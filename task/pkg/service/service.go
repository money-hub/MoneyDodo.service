package service

import (
	"context"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// TaskService describes the service.
type TaskService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	GetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	GetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	GetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
}

type basicTaskService struct {
	*db.DBService
}

func (b *basicTaskService) GetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of GetHisReleasedTasks
	role := ctx.Value("role").(int)
	if role == 1 {
		userID := ctx.Value("id").(string)
		if userID == id {
			task := model.Task{}
			rows, err := b.Engine().Where("Publisher = ?", id).Rows(task)
			if err == nil {
				for rows.Next() {
					err1 := rows.Scan(&task)
					if err1 != nil {
						return false, err1.Error(), data
					}
					data = append(data, task)
				}
				return true, "", data
			}
			return false, err.Error(), data
		}
	} else if role == 0 {
		task := model.Task{}
		rows, err := b.Engine().Where("Publisher = ? and state = ?", id, "released").Rows(task)
		if err == nil {
			for rows.Next() {
				err1 := rows.Scan(&task)
				if err1 != nil {
					return false, err1.Error(), data
				}
				data = append(data, task)
			}
			return true, "", data
		}
		return false, err.Error(), data
	}
	return false, "Permission denied", data
}
func (b *basicTaskService) GetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of GetTasksByID
	task := model.Task{}
	rows, err := b.Engine().Where("Publisher = ? and State = ?", id, "released").Rows(task)
	if err == nil {
		for rows.Next() {
			err1 := rows.Scan(&task)
			if err1 != nil {
				return false, err1.Error(), data
			}
			data = append(data, task)
		}
		return true, "", data
	}
	return false, err.Error(), data
}
func (b *basicTaskService) GetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of GetHisUnreleasedTasks
	role := ctx.Value("role").(int)
	if role == 1 {
		userID := ctx.Value("id").(string)
		if userID == id {
			task := model.Task{}
			rows, err := b.Engine().Where("Publisher = ? and State = ?", id, "non-released").Rows(task)
			if err == nil {
				for rows.Next() {
					err1 := rows.Scan(&task)
					if err1 != nil {
						return false, err1.Error(), data
					}
					data = append(data, task)
				}
				return true, "", data
			}
			return false, err.Error(), data
		}
	} else if role == 0 {
		task := model.Task{}
		rows, err := b.Engine().Where("Publisher = ? and state = ?", id, "non-released").Rows(task)
		if err == nil {
			for rows.Next() {
				err1 := rows.Scan(&task)
				if err1 != nil {
					return false, err1.Error(), data
				}
				data = append(data, task)
			}
			return true, "", data
		}
		return false, err.Error(), data
	}
	return false, "Permission denied", data
}
func (b *basicTaskService) GetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of GetHisClosedTasks
	role := ctx.Value("role").(int)
	if role == 1 {
		userID := ctx.Value("id").(string)
		if userID == id {
			task := model.Task{}
			rows, err := b.Engine().Where("Publisher = ? and State = ?", id, "closed").Rows(task)
			if err == nil {
				for rows.Next() {
					err1 := rows.Scan(&task)
					if err1 != nil {
						return false, err1.Error(), data
					}
					data = append(data, task)
				}
				return true, "", data
			}
			return false, err.Error(), data
		}
	} else if role == 0 {
		task := model.Task{}
		rows, err := b.Engine().Where("Publisher = ? and state = ?", id, "closed").Rows(task)
		if err == nil {
			for rows.Next() {
				err1 := rows.Scan(&task)
				if err1 != nil {
					return false, err1.Error(), data
				}
				data = append(data, task)
			}
			return true, "", data
		}
		return false, err.Error(), data
	}
	return false, "Permission denied", data
}

// NewBasicTaskService returns a naive, stateless implementation of TaskService.
func NewBasicTaskService() TaskService {
	basicTaskService := &basicTaskService{
		&db.DBService{},
	}
	err := basicTaskService.Bind("conf/conf.moneydodo.yml")
	if err != nil {
		log.Printf("The UserService failed to bind with mysql")
	}
	return basicTaskService
}

// New returns a TaskService with all of the expected middleware wired in.
func New(middleware []Middleware) TaskService {
	var svc TaskService = NewBasicTaskService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
