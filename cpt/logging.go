package cpt

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/money-hub/MoneyDodo.service/model"
)

// log middleware
type LoggingMiddleware struct {
	Logger log.Logger
	Next   CptService
}

func (l *LoggingMiddleware) GetSpec(ctx context.Context, taskId string) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetSpec",
			"input", fmt.Sprintf("%#v", Request{TaskId: taskId}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetSpec(ctx, taskId)
	return
}

func (l *LoggingMiddleware) GetAll(ctx context.Context, kind, state string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetAll",
			"input", fmt.Sprintf("%#v", Request{Kind: kind, State: state}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetAll(ctx, kind, state, page, offset, limit, orderby)
	return
}

func (l *LoggingMiddleware) Post(ctx context.Context, kind string, task interface{}) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Post",
			"input", fmt.Sprintf("%#v", Request{Kind: kind, Task: task}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Post(ctx, kind, task)
	return
}

func (l *LoggingMiddleware) Put(ctx context.Context, taskId string, task interface{}) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Put",
			"input", fmt.Sprintf("%#v", Request{TaskId: taskId, Task: task}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Put(ctx, taskId, task)
	return
}

func (l *LoggingMiddleware) Delete(ctx context.Context, taskId string, state string) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Delete",
			"input", fmt.Sprintf("%#v", Request{TaskId: taskId, State: state}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Delete(ctx, taskId, state)
	return
}
