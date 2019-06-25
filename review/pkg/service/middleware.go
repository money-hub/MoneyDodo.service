package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// Middleware describes a service middleware.
type Middleware func(ReviewService) ReviewService

type loggingMiddleware struct {
	logger log.Logger
	next   ReviewService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ReviewService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ReviewService) ReviewService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) PostReview(ctx context.Context, review model.Review) (status bool, errinfo string, data model.Review) {
	defer func() {
		l.logger.Log("method", "PostReview", "review", review, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.PostReview(ctx, review)
}
func (l loggingMiddleware) GetReview(ctx context.Context, rid string) (status bool, errinfo string, data model.Review) {
	defer func() {
		l.logger.Log("method", "GetReview", "rid", rid, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetReview(ctx, rid)
}
func (l loggingMiddleware) GetReviews(ctx context.Context) (status bool, errinfo string, data []model.Review) {
	defer func() {
		l.logger.Log("method", "GetReviews", "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetReviews(ctx)
}
func (l loggingMiddleware) PutReview(ctx context.Context, rid string, review model.Review) (status bool, errinfo string, data model.Review) {
	defer func() {
		l.logger.Log("method", "PutReview", "rid", rid, "review", review, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.PutReview(ctx, rid, review)
}
