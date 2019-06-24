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

func (l loggingMiddleware) GetTasksByID(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Task) {
	defer func() {
		l.logger.Log("method", "GetTasksByID", "id", id, "state", state, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetTasksByID(ctx, id, state)
}
