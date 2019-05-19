package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	"github.com/money-hub/MoneyDodo.service/model"
)

// Middleware describes a service middleware.
type Middleware func(TaskService) TaskService

type loggingMiddleware struct {
	logger log.Logger
	next   TaskService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TaskService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TaskService) TaskService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) UserGetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "UserGetHisReleasedTasks", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.UserGetHisReleasedTasks(ctx, id)
}
func (l loggingMiddleware) UserGetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "UserGetTasksByID", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.UserGetTasksByID(ctx, id)
}
func (l loggingMiddleware) UserGetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "UserGetHisUnreleasedTasks", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.UserGetHisUnreleasedTasks(ctx, id)
}
func (l loggingMiddleware) UserGetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "UserGetHisClosedTasks", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.UserGetHisClosedTasks(ctx, id)
}
func (l loggingMiddleware) AdminGetAllTasksByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "AdminGetAllTasksByUserID", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.AdminGetAllTasksByUserID(ctx, id)
}
func (l loggingMiddleware) AdminGetTasksReleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "AdminGetTasksReleasedByUserID", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.AdminGetTasksReleasedByUserID(ctx, id)
}
func (l loggingMiddleware) AdminGetTasksUnreleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "AdminGetTasksUnreleasedByUserID", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.AdminGetTasksUnreleasedByUserID(ctx, id)
}
func (l loggingMiddleware) AdminGetTasksClosedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "AdminGetTasksClosedByUserID", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.AdminGetTasksClosedByUserID(ctx, id)
}
