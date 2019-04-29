package cpt

import (
	"context"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/money-hub/MoneyDodo.service/model"
)

// log middleware
type LoggingMiddleware struct {
	Logger log.Logger
	Next   CptService
}

func (l *LoggingMiddleware) GetSpec(ctx context.Context, taskId string) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetSpec",
			"input", Request{TaskId: taskId},
			"output", Response{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetSpec(ctx, taskId)
	return
}

func (l *LoggingMiddleware) GetAll(ctx context.Context, state string) (status bool, errinfo string, data []model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetAll",
			"input", Request{State: state},
			"output", Response{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetAll(ctx, state)
	return
}

func (l *LoggingMiddleware) Post(ctx context.Context, task model.Task) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Post",
			"input", Request{Task: task},
			"output", Response{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Post(ctx, task)
	return
}

func (l *LoggingMiddleware) Put(ctx context.Context, taskId string, action string, task model.Task) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Put",
			"input", Request{TaskId: taskId, Action: action, Task: task},
			"output", Response{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Put(ctx, taskId, action, task)
	return
}

func (l *LoggingMiddleware) Delete(ctx context.Context, taskId string, state string) (status1 bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Delete",
			"input", Request{TaskId: taskId, State: state},
			"output", Response{status1, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status1, errinfo, data = l.Next.Delete(ctx, taskId, state)
	return
}
