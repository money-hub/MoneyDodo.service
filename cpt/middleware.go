package personalTasks

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

func (l *LoggingMiddleware) GetSpec(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetSpec",
			"input", GetSpecRequest{userId, taskId},
			"output", GetSpecResponse{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetSpec(ctx, userId, taskId)
	return
}

func (l *LoggingMiddleware) GetAll(ctx context.Context, userId string) (status bool, errinfo string, data []model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetAll",
			"input", GetAllRequest{userId},
			"output", GetAllResponse{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetAll(ctx, userId)
	return
}

func (l *LoggingMiddleware) Post(ctx context.Context, userId string, task model.Task) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Post",
			"input", PostRequest{userId, task},
			"output", PostResponse{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Post(ctx, userId, task)
	return
}

func (l *LoggingMiddleware) PostClaim(ctx context.Context, userId string, taskId string) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "PostClaim",
			"input", PostClaimRequest{userId, taskId},
			"output", PostClaimResponse{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.PostClaim(ctx, userId, taskId)
	return
}

func (l *LoggingMiddleware) Put(ctx context.Context, userId string, taskId string, task model.Task) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Put",
			"input", PutRequest{userId, taskId, task},
			"output", PutResponse{status, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Put(ctx, userId, taskId, task)
	return
}

func (l *LoggingMiddleware) Delete(ctx context.Context, userId string, taskId string, status string) (status1 bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Delete",
			"input", DeleteRequest{userId, taskId, status},
			"output", DeleteResponse{status1, errinfo, data},
			"took", time.Since(begin),
		)
	}(time.Now())
	status1, errinfo, data = l.Next.Delete(ctx, userId, taskId, status)
	return
}
