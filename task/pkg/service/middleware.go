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

func (l loggingMiddleware) GetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "GetHisReleasedTasks", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetHisReleasedTasks(ctx, id)
}
func (l loggingMiddleware) GetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "GetTasksByID", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetTasksByID(ctx, id)
}
func (l loggingMiddleware) GetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "GetHisUnreleasedTasks", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetHisUnreleasedTasks(ctx, id)
}
func (l loggingMiddleware) GetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "GetHisClosedTasks", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetHisClosedTasks(ctx, id)
}
