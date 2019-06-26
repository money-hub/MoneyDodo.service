package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// Middleware describes a service middleware.
type Middleware func(CommentService) CommentService

type loggingMiddleware struct {
	logger log.Logger
	next   CommentService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a CommentService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next CommentService) CommentService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetComment(ctx context.Context, taskId string) (status bool, errinfo string, data []model.Comment) {
	defer func() {
		l.logger.Log("method", "GetComment", "taskId", taskId, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetComment(ctx, taskId)
}
func (l loggingMiddleware) PostComment(ctx context.Context, taskId string, comment string) (status bool, errinfo string, data *model.Comment) {
	defer func() {
		l.logger.Log("method", "PostComment", "taskId", taskId, "comment", comment, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.PostComment(ctx, taskId, comment)
}
func (l loggingMiddleware) ChangeComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data *model.Comment) {
	defer func() {
		l.logger.Log("method", "ChangeComment", "taskId", taskId, "cId", cId, "comment", comment, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.ChangeComment(ctx, taskId, cId, comment)
}
func (l loggingMiddleware) DeleteComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data string) {
	defer func() {
		l.logger.Log("method", "DeleteComment", "taskId", taskId, "cId", cId, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.DeleteComment(ctx, taskId, cId)
}
func (l loggingMiddleware) LikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment) {
	defer func() {
		l.logger.Log("method", "LikeComment", "taskId", taskId, "cId", cId, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.LikeComment(ctx, taskId, cId)
}
func (l loggingMiddleware) CancelLikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment) {
	defer func() {
		l.logger.Log("method", "CancelLikeComment", "taskId", taskId, "cId", cId, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.CancelLikeComment(ctx, taskId, cId)
}
