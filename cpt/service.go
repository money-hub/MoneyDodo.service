package cpt

import (
	"context"
	"log"
	"time"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// BASE_URL=http://hostname:port/api/tasks

type CptService interface {
	// 查询所有任务、未发布的任务、已发布但未被领取、正在进行、已经完成
	GetAll(ctx context.Context, state string) (status bool, errinfo string, data []model.Task)
	// 查询某个任务
	GetSpec(ctx context.Context, taskId string) (status bool, errinfo string, data *model.Task)
	// 创建任务
	Post(ctx context.Context, task model.Task) (status bool, errinfo string, data *model.Task)
	// 发布任务、领取任务、完成任务
	Put(ctx context.Context, taskId string, action string, task model.Task) (status bool, errinfo string, data *model.Task)
	// 取消领取任务、删除已完成任务、取消发布任务
	Delete(ctx context.Context, taskId string, state string) (status bool, errinfo string, data *model.Task)
}

type basicCptService struct {
	*db.DBService
}

func (b *basicCptService) GetAll(ctx context.Context, state string) (status bool, errinfo string, data []model.Task) {
	data = make([]model.Task, 0)
	var err error
	if state == "" {
		err = b.Engine().Find(&data)
	} else {
		err = b.Engine().Where("state = ?", state).Find(&data)
	}
	status = err == nil
	if err != nil {
		errinfo = err.Error()
	}
	return
}

func (b *basicCptService) GetSpec(ctx context.Context, taskId string) (status bool, errinfo string, data *model.Task) {
	data = &model.Task{
		Id: taskId,
	}
	status, err := b.Engine().Get(data)
	if status == false {
		data = nil
	}
	if err != nil {
		errinfo = err.Error()
	}
	return
}

func (b *basicCptService) Post(ctx context.Context, task model.Task) (status bool, errinfo string, data *model.Task) {
	var err error
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}
	if _, err = sess.Insert(task); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	lastTask := &model.Task{}
	if _, err = sess.Limit(1, 0).Desc("Id").Get(lastTask); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", lastTask
}

func (b *basicCptService) Put(ctx context.Context, taskId string, action string, task model.Task) (status bool, errinfo string, data *model.Task) {
	// 判断当前任务是否存在
	task1 := &model.Task{
		Id: taskId,
	}
	ok, err := b.Engine().Get(task1)
	if err != nil {
		return false, err.Error(), nil
	}
	if !ok {
		return false, "The task doesn't existed.", nil
	}

	// 判断action与task.State是否一致
	if task.State != action {
		return false, "The task.State is not equal to action", nil
	}

	// 获取进行此操作的用户ID
	id := ctx.Value("Id").(string)

	// 进行具体的action
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}
	// 任务创建者
	if task1.Publisher == id {
		// 1. 如果当前操作为release，则需要满足之前task的state为non-released
		// 2. 如果当前操作为finish，则需要满足之前task的state为claimed
		if (action == model.TaskActionRelease && task1.State == model.TaskStateNonReleased) || (action == model.TaskActionFinish && task1.State == model.TaskStateClaimed) {
			if action == model.TaskActionRelease {
				task.Pubdate = time.Now()
			} else if action == model.TaskActionFinish {
				task.Enddate = time.Now()
				task.ConfirmFinish = true
			}
			_, err := sess.Where("id = ?", taskId).Update(task)
			if err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		} else {
			return false, "The url parameter action is not matching to task.State.", nil
		}
	} else if task1.Recipient == id {
		// 任务接受者
		// 1. 如果当前操作为claim，则需要满足之前task的state为released
		// 2. 如果当前操作为finish，则需要满足之前task的state为claimed
		if (action == model.TaskActionClaim && task1.State == model.TaskStateReleased) || (action == model.TaskActionFinish && task1.State == model.TaskStateClaimed) {
			if action == model.TaskActionFinish {
				task.RecipientFinish = true
			}
			_, err := sess.Where("id = ?", taskId).Update(task)
			if err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			if action == model.TaskActionClaim {
				deal := model.Deal{
					TaskId:    taskId,
					Publisher: task.Publisher,
					Recipient: task.Recipient,
					Since:     time.Now(),
					Reward:    task.Reward,
					State:     model.DealStateUnderway,
				}
				if _, err = sess.Insert(deal); err != nil {
					sess.Rollback()
					return false, err.Error(), nil
				}
			}
		} else {
			return false, "The url parameter action is not matching to task.State.", nil
		}
	} else {
		return false, "The user is not involved in this task.", nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", nil
}

func (b *basicCptService) Delete(ctx context.Context, taskId string, state string) (status bool, errinfo string, data *model.Task) {
	if state != model.TaskStateNonReleased && state != model.TaskStateReleased && state != model.TaskStateClaimed && state != model.TaskStateFinished {
		return false, "The query parameter status is not correct!", nil
	}
	var err error
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}
	deal := model.Deal{
		TaskId: taskId,
	}

	task := &model.Task{
		Id: taskId,
	}
	if _, err = sess.Get(task); err != nil {
		return false, err.Error(), nil
	}

	// 获取进行此操作的用户ID
	id := ctx.Value("Id").(string)
	// 发布者
	if task.Publisher == id {
		if state == model.TaskStateClaimed {
			return false, "The query parameter state is not correct!", nil
		}
		if state == model.TaskStateReleased {
			if task.State == model.TaskStateNonReleased {
				return true, "The task doesn't released, you don't need to cancel it.", nil
			}
			// 发布者不能取消一个被接受的任务
			if task.State == model.TaskStateClaimed {
				return false, "The task has been claimed, you can't cancel it.", nil
			}

			if task.State == model.TaskStateFinished {
				return false, "The task has been finished, you can't cancel it", nil
			}

			// 发布者取消发布，更新任务状态为未发布
			task.State = model.TaskStateNonReleased
			if _, err = sess.Where("Id = ?", taskId).Update(task); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}

		// 发布者删除未发布的任务
		if state == model.TaskStateNonReleased {
			if _, err = sess.Delete(task); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}
		// 发布者删除已完成的任务
		if state == model.TaskStateFinished {
			// 如果接收者已经删除了此任务，则将任务从数据库中删除
			if task.Recipient != "" {
				if _, err = sess.Delete(task); err != nil {
					sess.Rollback()
					return false, err.Error(), nil
				}
			} else {
				// 否则，只需要将task.Publisher置为空即可
				task.Publisher = ""
				if _, err = sess.Where("Id = ?", taskId).AllCols().Update(task); err != nil {
					sess.Rollback()
					return false, err.Error(), nil
				}
			}
		}
	} else if task.Recipient == id {
		// 接收者
		// 取消接受任务，删除deal交易信息，更新任务状态为为发布状态
		if state == model.TaskStateClaimed {
			if _, err = sess.Delete(deal); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			task.State = model.TaskStateReleased
			if _, err = sess.Where("Id = ?", taskId).AllCols().Update(task); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		} else if state == model.TaskStateFinished {
			// 如果发布者已经删除了此任务，则将任务从数据库中删除
			if task.Publisher != "" {
				if _, err = sess.Delete(task); err != nil {
					sess.Rollback()
					return false, err.Error(), nil
				}
			} else {
				// 否则，只需要将task.Recipient置为空即可
				task.Recipient = ""
				if _, err = sess.Where("Id = ?", taskId).AllCols().Update(task); err != nil {
					sess.Rollback()
					return false, err.Error(), nil
				}
			}
		} else {
			return false, "The query parameter status is not correct!", nil
		}
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", nil
}

// NewBasicCptService returns a naive, stateless implementation of CptService.
func NewBasicCptService() CptService {
	basicCptSvc := &basicCptService{
		&db.DBService{},
	}
	err := basicCptSvc.Bind("conf/conf.lyh.yml")
	if err != nil {
		log.Printf("The CptService failed to bind with mysql")
	}
	return basicCptSvc
}

// // New returns a CptService with all of the expected middleware wired in.
// func New(middleware []Middleware) CptService {
// 	var svc CptService = NewBasicCptService()
// 	for _, m := range middleware {
// 		svc = m(svc)
// 	}
// 	return svc
// }
