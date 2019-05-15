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
	UserGetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	UserGetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	UserGetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	UserGetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	AdminGetAllTasksByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	AdminGetTasksReleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	AdminGetTasksUnreleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
	AdminGetTasksClosedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task)
}

type basicTaskService struct {
	*db.DBService
}

func (b *basicTaskService) UserGetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of UserGetHisReleasedTasks
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
	}
	return false, "Permission denied", data
}
func (b *basicTaskService) UserGetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of UserGetTasksByID
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
func (b *basicTaskService) UserGetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of UserGetHisUnreleasedTasks
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
	}
	return false, "Permission denied", data
}
func (b *basicTaskService) UserGetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of UserGetHisClosedTasks
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
	}
	return false, "Permission denied", data
}
func (b *basicTaskService) AdminGetAllTasksByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of AdminGetAllTasksByUserID
	role := ctx.Value("role").(int)
	if role == 0 {
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
	return false, "Permission denied", data
}
func (b *basicTaskService) AdminGetTasksReleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of AdminGetTasksReleasedByUserID
	role := ctx.Value("role").(int)
	if role == 0 {
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
func (b *basicTaskService) AdminGetTasksUnreleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of AdminGetTasksUnreleasedByUserID
	role := ctx.Value("role").(int)
	if role == 0 {
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
func (b *basicTaskService) AdminGetTasksClosedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	// TODO implement the business logic of AdminGetTasksClosedByUserID
	role := ctx.Value("role").(int)
	if role == 0 {
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
	var svc = NewBasicTaskService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
