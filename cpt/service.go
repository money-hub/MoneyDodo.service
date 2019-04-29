package personalTasks

import (
	"context"
	"log"

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
		if (action == model.TaskActionRelease && task1.State == model.TaskStateNonReleased) || (action == model.TaskActionFinish && task1.State == model.TaskActionClaim) {
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
		if action == model.TaskActionClaim && task1.State == model.TaskStateReleased {
			_, err := sess.Where("id = ?", taskId).Update(task)
			if err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			deal := model.Deal{
				TaskId: taskId,
			}
			if _, err = sess.Where("userId = ? and taskId = ?", userId, taskId).AllCols().Update(r); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		} else {
			return false, "The url parameter action is not matching to task.State.", nil
		}
	} else {
		return false, "You are not involved in this task.", nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", nil
}

func (b *basicCptService) Delete(ctx context.Context, userId string, taskId string, status string) (status1 bool, errinfo string, data *model.Task) {
	// 用户未注册
	user := &model.User{
		Id: userId,
	}
	ok, err := b.Engine().Get(user)
	if err != nil {
		return false, err.Error(), nil
	}
	if !ok {
		return false, "The user may not be registered.", nil
	}
	if status != model.TaskStatusNone && status != model.TaskStatusReleased && status != model.TaskStatusClaimed && status != model.TaskStatusFinished {
		return false, "The query parameter status is not correct!", nil
	}
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}
	r := model.Relation{
		UserId: userId,
		TaskId: taskId,
		Status: status,
	}

	task := &model.Task{
		Id: taskId,
	}
	if _, err = sess.Get(task); err != nil {
		return false, err.Error(), nil
	}

	task1 := *task
	task1.Recipient = ""
	task1.Status = model.TaskStatusNone

	// 发布者
	if task.Publisher == userId {
		log.Println(1)
		if status == model.TaskStatusClaimed {
			return false, "The query parameter status is not correct!", nil
		}
		if status == model.TaskStatusReleased {
			if task.Status == model.TaskStatusNone {
				return true, "The task doesn't released, you don't need to cancel it.", nil
			}
			// 发布者不能取消一个被接受的任务
			if task.Status == model.TaskStatusClaimed {
				return false, "The task has been claimed, you can't cancel it.", nil
			}

			if task.Status == model.TaskStatusFinished {
				return false, "The task has been finished, you can't cancel it", nil
			}

			// 发布者取消发布，或者删除已完成的任务
			if _, err = sess.Delete(r); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			// 更新任务状态为未发布
			if _, err = sess.Where("Id = ?", taskId).AllCols().Update(task1); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}

		if status == model.TaskStatusNone || status == model.TaskStatusFinished {
			if _, err = sess.Delete(task); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}
	} else if task.Recipient == userId {
		// 领取者
		log.Println(2)
		if status != model.TaskStatusClaimed {
			return false, "The query parameter status is not correct!", nil
		}
		if _, err = sess.Delete(r); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
		// 更新任务状态为未发布
		if _, err = sess.Where("Id = ?", taskId).AllCols().Update(task1); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return false, "The url userId is not equal to the task.Publisher or task.Recipient.", nil
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
