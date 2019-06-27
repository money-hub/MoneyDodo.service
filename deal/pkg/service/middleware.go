package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// Middleware describes a service middleware.
type Middleware func(DealService) DealService

type loggingMiddleware struct {
	logger log.Logger
	next   DealService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a DealService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next DealService) DealService {
		return &loggingMiddleware{logger, next}
	}

}
func (l loggingMiddleware) PutDealState(ctx context.Context, deal model.Deal) (status bool, errinfo string, data *model.Deal) {
	defer func() {
		l.logger.Log("method", "PutDealState", "deal", deal, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.PutDealState(ctx, deal)
}
func (l loggingMiddleware) GetUserDealByState(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Deal) {
	defer func() {
		l.logger.Log("method", "GetUserDealByState", "id", id, "state", state, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetUserDealByState(ctx, id, state)
}
func (l loggingMiddleware) GetDealByDID(ctx context.Context, id string) (status bool, errinfo string, data model.Deal) {
	defer func() {
		l.logger.Log("method", "GetDealByDID", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetDealByDID(ctx, id)
}
func (l loggingMiddleware) GetDealByState(ctx context.Context, state string) (status bool, errinfo string, data []model.Deal) {
	defer func() {
		l.logger.Log("method", "GetDealByState", "state", state, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetDealByState(ctx, state)
}
func (l loggingMiddleware) PostAcceptDeal(ctx context.Context, deal model.Deal) (status bool, errinfo string, data model.Deal) {
	defer func() {
		l.logger.Log("method", "PostAcceptDeal", "deal", deal, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.PostAcceptDeal(ctx, deal)
}
