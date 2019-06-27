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
	GetTasksByID(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Task)
}

type basicTaskService struct {
	*db.DBService
}

func (b *basicTaskService) GetTasksByID(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of GetTasksByID
	task := model.Task{}
	role := ctx.Value("role").(int)
	userID := ctx.Value("id").(string)
	if state == "released" {
		rows, err := b.Engine().Where("publisher = ? and state = ?", id, state).Rows(task)
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
	} else if state == "" {
		if role == 1 || role == 2 {
			if userID == id {
				rows, err := b.Engine().Where("publisher = ?", id).Rows(task)
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
		} else {
			rows, err := b.Engine().Where("publisher = ?", id).Rows(task)
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
	} else {
		if role == 1 || role == 2 {
			if userID == id {
				rows, err := b.Engine().Where("publisher = ? and state = ?", id, state).Rows(task)
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
		} else {
			rows, err := b.Engine().Where("publisher = ? and state = ?", id, state).Rows(task)
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
