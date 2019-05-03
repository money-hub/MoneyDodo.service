package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetSpec(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "GetSpec", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetSpec(ctx, id)
}
func (l loggingMiddleware) GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.User) {
	defer func() {
		l.logger.Log("method", "GetAll", "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetAll(ctx, page, offset, limit, orderby)
}
func (l loggingMiddleware) GetUDF(ctx context.Context, name string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.User) {
	defer func() {
		l.logger.Log("method", "GetUDF", "name", name, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetUDF(ctx, name, page, offset, limit, orderby)
}
func (l loggingMiddleware) Post(ctx context.Context, user model.User) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "Post", "user", user, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.Post(ctx, user)
}
func (l loggingMiddleware) Patch(ctx context.Context, id string, user model.User) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "Patch", "id", id, "user", user, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.Patch(ctx, id, user)
}
func (l loggingMiddleware) Put(ctx context.Context, id string, user model.User) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "Put", "id", id, "user", user, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.Put(ctx, id, user)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.Delete(ctx, id)
}
