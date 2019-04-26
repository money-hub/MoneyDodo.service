package personalTasks

import (
	"context"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

type PTaskService interface {
	GetAll(ctx context.Context, userId string) (status bool, errinfo string, data []model.Task)
	GetSpec(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task)
	Post(ctx context.Context, userId string, task model.Task) (status bool, errinfo string, data *model.Task)
	PostClaim(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task)
	Put(ctx context.Context, userId string, taskId string, task model.Task) (status bool, errinfo string, data *model.Task)
	Delete(ctx context.Context, userId string, taskId string, status string) (status1 bool, errinfo string, data *model.Task)
}

type basicPTaskService struct {
	*db.DBService
}

func (b *basicPTaskService) GetSpec(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task) {
	// TODO implement the business logic of GetSpec
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
	data = &model.Task{
		Id: taskId,
	}
	status, err = b.Engine().Get(data)
	if data.Publisher != userId {
		return false, "The user may not have release this task. ", nil
	}
	if status == false {
		data = nil
	}
	if err != nil {
		errinfo = err.Error()
	}
	return
}

func (b *basicPTaskService) GetAll(ctx context.Context, userId string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of GetAll
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
	data = make([]model.Task, 0)
	err = b.Engine().Where("publisher = ?", userId).Find(&data)
	status = err == nil
	if err != nil {
		errinfo = err.Error()
	}
	return
}

func (b *basicPTaskService) Post(ctx context.Context, userId string, task model.Task) (status bool, errinfo string, data *model.Task) {
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

	if userId != task.Publisher {
		return false, "The url userId is not equal to the task.Publisher.", nil
	}
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

func (b *basicPTaskService) PostClaim(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task) {
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
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}

	t := &model.Task{
		Id: taskId,
	}

	if _, err := b.Engine().Get(t); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	if t.Publisher == userId {
		sess.Rollback()
		return false, "The user can't claim his/her own published tasks.", nil
	}

	// 查询该task是否存在或被其他用户已经领取
	task1 := &model.Task{
		Id: taskId,
	}

	ok, err = b.Engine().Get(task1)
	if err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	if !ok {
		sess.Rollback()
		return false, "The task doesn't existed.", nil
	}

	if task1.Status != model.TaskStatusReleased {
		sess.Rollback()
		return false, "The task doesn't release or has been claimed by others.", nil
	}

	task := model.Task{
		Id:        taskId,
		Recipient: userId,
		Status:    model.TaskStatusClaimed,
	}
	if _, err := b.Engine().Where("id = ?", taskId).Update(task); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	r := model.Relation{
		UserId: userId,
		TaskId: taskId,
		Status: model.TaskStatusClaimed,
	}
	if _, err := b.Engine().Insert(r); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", nil
}

func (b *basicPTaskService) Put(ctx context.Context, userId string, taskId string, task model.Task) (status bool, errinfo string, data *model.Task) {
	// userId != publisher
	if userId != task.Publisher {
		return false, "The url userId is not equal to the task.Publisher.", nil
	}
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

	task1 := &model.Task{
		Id: taskId,
	}

	ok, err = b.Engine().Get(task1)
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
	// 发布者
	if task1.Publisher == userId {
		if (task1.Status == model.TaskStatusNone && task.Status == model.TaskStatusNone) || task1.Status == model.TaskStatusNone && task.Status == model.TaskStatusReleased {
			_, err := sess.Where("id = ? and publisher = ?", taskId, userId).AllCols().Update(task)
			if err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			r := model.Relation{
				UserId: userId,
				TaskId: taskId,
				Status: model.TaskStatusReleased,
			}
			if _, err = sess.Insert(r); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}
	} else if task1.Recipient == userId {
		// 接受者
		if (task1.Status == model.TaskStatusReleased && task.Status == model.TaskStatusClaimed) || (task1.Status == model.TaskStatusClaimed && task.Status == model.TaskStatusFinished) {
			_, err := sess.Where("id = ? and recipient = ?", taskId, userId).AllCols().Update(task)
			if err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			r := model.Relation{
				UserId: userId,
				TaskId: taskId,
				Status: task.Status,
			}
			if _, err = sess.Where("userId = ? and taskId = ?", userId, taskId).AllCols().Update(r); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return false, "The url userId is not equal to the task.Publisher or task.Recipient.", nil
}

func (b *basicPTaskService) Delete(ctx context.Context, userId string, taskId string, status string) (status1 bool, errinfo string, data *model.Task) {
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

// NewBasicPTaskService returns a naive, stateless implementation of PTaskService.
func NewBasicPTaskService() PTaskService {
	basicPTaskSvc := &basicPTaskService{
		&db.DBService{},
	}
	err := basicPTaskSvc.Bind("conf/conf.lyh.yml")
	if err != nil {
		log.Printf("The PTaskService failed to bind with mysql")
	}
	return basicPTaskSvc
}

// // New returns a PTaskService with all of the expected middleware wired in.
// func New(middleware []Middleware) PTaskService {
// 	var svc PTaskService = NewBasicPTaskService()
// 	for _, m := range middleware {
// 		svc = m(svc)
// 	}
// 	return svc
// }
