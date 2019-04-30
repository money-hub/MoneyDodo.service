package cpt

import (
	"context"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// BASE_URL=http://hostname:port/api/tasks

type CptService interface {
	// 查询所有任务，kind和state为可选参数，返回值为缩略信息
	GetAll(ctx context.Context, kind string, state string) (status bool, errinfo string, data []model.Task)
	// 查询某个任务，返回值为相应类型的详细信息
	GetSpec(ctx context.Context, taskId string) (status bool, errinfo string, data interface{})
	// 创建任务，上传任务为详细信息
	Post(ctx context.Context, kind string, task interface{}) (status bool, errinfo string, data interface{})
	// 发布任务、领取任务、完成任务
	Put(ctx context.Context, taskId string, task interface{}) (status bool, errinfo string, data interface{})
	// // 取消领取任务、删除已完成任务、取消发布任务
	Delete(ctx context.Context, taskId string, state string) (status bool, errinfo string, data *model.Task)
}

type basicCptService struct {
	*db.DBService
}

func (b *basicCptService) GetAll(ctx context.Context, kind string, state string) (status bool, errinfo string, data []model.Task) {
	data = make([]model.Task, 0)
	var err error
	if kind == "" && state == "" {
		err = b.Engine().Find(&data)
	} else if kind == "" {
		err = b.Engine().Where("state = ?", state).Find(&data)
	} else if state == "" {
		err = b.Engine().Where("kind = ?", kind).Find(&data)
	} else {
		err = b.Engine().Where("kind = ? and state = ?", kind, state).Find(&data)
	}
	status = err == nil
	if err != nil {
		errinfo = err.Error()
	}
	return
}

func (b *basicCptService) GetSpec(ctx context.Context, taskId string) (status bool, errinfo string, data interface{}) {
	var err error
	var ok bool
	task := model.Task{
		Id: taskId,
	}
	status, err = b.Engine().Get(&task)
	if status == false {
		return false, "The query task is not existed.", nil
	}
	if err != nil {
		return false, err.Error(), nil
	}
	if task.Kind == model.TaskKindQuestionnaire {
		qtnr := model.Questionnaire{
			TaskId: taskId,
		}
		if ok, err = b.Engine().Get(&qtnr); err != nil {
			return false, err.Error(), nil
		}
		if !ok {
			return false, "The questionnaire has not been created.", nil
		}
		data = model.Qtnr{
			Task: task,
			Qtnr: &qtnr,
		}
		return true, "", data
	} else {
		return false, "The task kind is not true", nil
	}
}

func (b *basicCptService) Post(ctx context.Context, kind string, taskUnknown interface{}) (status bool, errinfo string, data interface{}) {
	var err error
	if kind == model.TaskKindQuestionnaire {
		qtnr, ok := taskUnknown.(model.Qtnr)
		if !ok {
			return false, "The task kind is not matching to the uploaded task.", nil
		}
		sess := b.Engine().NewSession()
		defer sess.Close()
		if err = sess.Begin(); err != nil {
			return false, err.Error(), nil
		}
		qtnr.State = model.TaskStateNonReleased
		if _, err = sess.Insert(qtnr.Task); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
		lastTask := &model.Task{}
		if _, err = sess.Limit(1, 0).Desc("Id").Get(lastTask); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
		qtnr.Id = lastTask.Id
		qtnr.Qtnr.TaskId = lastTask.Id
		if _, err = sess.Insert(qtnr.Qtnr); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
		err = sess.Commit()
		if err != nil {
			return false, err.Error(), nil
		}
		return true, "", qtnr
	} else {
		return false, "The task kind is not true", nil
	}
}

func (b *basicCptService) Put(ctx context.Context, taskId string, taskUnknown interface{}) (status bool, errinfo string, data interface{}) {
	// 判断当前任务是否存在
	task := &model.Task{
		Id: taskId,
	}
	ok, err := b.Engine().Get(task)
	if err != nil {
		return false, err.Error(), nil
	}
	if !ok {
		return false, "The task doesn't existed.", nil
	}
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}

	if task.Kind == model.TaskKindQuestionnaire {
		qtnr, ok := taskUnknown.(model.Qtnr)
		if !ok {
			return false, "The task kind is not matching to the uploaded task.", nil
		}
		// 判断当前用户是否是任务的发布者
		if qtnr.Publisher != ctx.Value("Id").(string) {
			return false, "You are not permitted to modify others' task.", nil
		}
		if qtnr.State == model.TaskStateReleased {
			return false, "You can't modify the task which has been released.", nil
		}
		if _, err := sess.Where("id = ?", qtnr.Id).AllCols().Update(qtnr.Task); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
		if _, err := sess.Where("taskId = ?", qtnr.Id).AllCols().Update(qtnr.Qtnr); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
	} else {
		return false, "The task kind is not true", nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", nil
}

func (b *basicCptService) Delete(ctx context.Context, taskId string, state string) (status bool, errinfo string, data *model.Task) {
	if state != model.TaskStateNonReleased && state != model.TaskStateReleased && state != model.TaskStateClosed {
		return false, "The query parameter status is not correct!", nil
	}
	var err error
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}

	task := &model.Task{
		Id: taskId,
	}
	if _, err = sess.Get(task); err != nil {
		return false, err.Error(), nil
	}
	deal := make([]model.Deal, 0)
	if err = sess.Where("TaskId = ?", taskId).Find(deal); err != nil {
		return false, err.Error(), nil
	}

	// 获取进行此操作的用户ID
	id := "16340157"
	if task.Publisher != id {
		return false, "You are not permitted to modify others' task.", nil
	}

	if state == model.TaskStateReleased {
		if task.State == model.TaskStateNonReleased {
			return true, "The task doesn't released, you don't need to cancel it.", nil
		}
		// 发布者不能取消一个被接受的任务
		if len(deal) > 0 {
			return false, "You can't cancel task that has been claimed.", nil
		}

		if task.State == model.TaskStateClosed {
			return false, "The task has been finished, you can't cancel it", nil
		}

		// 发布者取消发布，更新任务状态为未发布
		task.State = model.TaskStateNonReleased
		if _, err = sess.Where("Id = ?", taskId).Update(task); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
	}

	// 发布者删除未发布的任务或者删除已完成的任务
	if state == model.TaskStateNonReleased || state == model.TaskStateClosed {
		questionnaire := model.Questionnaire{
			TaskId: taskId,
		}
		if _, err = sess.Delete(questionnaire); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
		if _, err = sess.Delete(task); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", nil
}

// NewBasicCptService returns a naive, stateless implementation of CptService.
func NewBasicCptService(conf string) CptService {
	basicCptSvc := &basicCptService{
		&db.DBService{},
	}
	err := basicCptSvc.Bind(conf)
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
