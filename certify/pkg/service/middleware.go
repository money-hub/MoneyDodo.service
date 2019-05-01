package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// Middleware describes a service middleware.
type Middleware func(CertifyService) CertifyService

type loggingMiddleware struct {
	logger log.Logger
	next   CertifyService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a CertifyService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next CertifyService) CertifyService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetAuthInfo(ctx context.Context, id string) (status bool, errinfo string, data model.User) {
	defer func() {
		l.logger.Log("method", "GetAuthInfo", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetAuthInfo(ctx, id)
}
func (l loggingMiddleware) PostAuthInfo(ctx context.Context, id string, img string) (status bool, errinfo string, data model.User) {
	defer func() {
		l.logger.Log("method", "PostAuthInfo", "id", id, "img", img, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.PostAuthInfo(ctx, id, img)
}
func (l loggingMiddleware) GetAllUnCertify(ctx context.Context) (status bool, errinfo string, data []model.User) {
	defer func() {
		l.logger.Log("method", "GetAllUnCertify", "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetAllUnCertify(ctx)
}
func (l loggingMiddleware) GetUnCertifyInfo(ctx context.Context, id string) (status bool, errinfo string, data model.User) {
	defer func() {
		l.logger.Log("method", "GetUnCertifyInfo", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetUnCertifyInfo(ctx, id)
}
func (l loggingMiddleware) PostCertifyState(ctx context.Context, id string, pass bool) (status bool, errinfo string, data model.User) {
	defer func() {
		l.logger.Log("method", "PostCertifyState", "id", id, "pass", pass, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.PostCertifyState(ctx, id, pass)
}
