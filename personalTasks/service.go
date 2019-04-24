package personalTasks

import (
	"context"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

const (
	Release = "Release"
	Claim   = "Claim"
)

type PTaskService interface {
	GetAll(ctx context.Context, userId string) (status bool, errinfo string, data []model.Task)
	GetSpec(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task)
	Post(ctx context.Context, userId string, task model.Task) (status bool, errinfo string, data *model.Task)
	PostClaim(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task)
	Put(ctx context.Context, userId string, taskId string, task model.Task) (status bool, errinfo string, data *model.Task)
	Delete(ctx context.Context, userId string, taskId string, detail string) (status bool, errinfo string, data *model.Task)
}

type basicPTaskService struct {
	*db.DBService
}

func (b *basicPTaskService) GetSpec(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task) {
	// TODO implement the business logic of GetSpec
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

func (b *basicPTaskService) GetAll(ctx context.Context, userId string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of GetAll
	data = make([]model.Task, 0)
	err := b.Engine().Where("From = ?", userId).Find(&data)
	status = err == nil
	if err != nil {
		errinfo = err.Error()
	}
	return
}

func (b *basicPTaskService) Post(ctx context.Context, userId string, task model.Task) (status bool, errinfo string, data *model.Task) {
	var ok bool
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
	if ok, err = sess.Limit(1, 0).Desc("Id").Get(lastTask); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	if ok {
		r := model.NewRelation(userId, lastTask.Id, Release)
		if _, err = sess.Insert(r); err != nil {
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

func (b *basicPTaskService) PostClaim(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task) {
	task := model.Task{
		Id:        taskId,
		Recipient: userId,
	}
	row, err := b.Engine().Where("Id = ?", taskId).Update(task)
	if err != nil {
		errinfo = err.Error()
	}
	return row > 0, errinfo, nil
}

func (b *basicPTaskService) Put(ctx context.Context, userId string, taskId string, task model.Task) (status bool, errinfo string, data *model.Task) {
	row, err := b.Engine().Where("id = ?", task).AllCols().Update(task)
	if err != nil {
		errinfo = err.Error()
	}
	return row > 0, errinfo, nil
}

func (b *basicPTaskService) Delete(ctx context.Context, userId string, taskId string, detail string) (status bool, errinfo string, data *model.Task) {
	if detail != Release || detail != Claim {
		return false, "The query parameter detail is not correct!", nil
	}
	var err error
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}
	r := model.Relation{
		UserId: userId,
		TaskId: taskId,
		Detail: detail,
	}

	if _, err = sess.Delete(r); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}

	if detail == Release {
		task := model.Task{
			Id: taskId,
		}
		if _, err = sess.Delete(task); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
	} else if detail == Claim {
		task := model.Task{
			Recipient: "",
		}
		if _, err = sess.Where("Id = ?", taskId).AllCols().Update(task); err != nil {
			sess.Rollback()
			return false, err.Error(), nil
		}
	}
	return true, "", nil
}

// NewBasicPTaskService returns a naive, stateless implementation of PTaskService.
func NewBasicPTaskService() PTaskService {
	basicPTaskSvc := &basicPTaskService{
		&db.DBService{},
	}
	err := basicPTaskSvc.Bind("conf/conf.users.yml")
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
