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

func (l loggingMiddleware) GetSpec(ctx context.Context, id string) (b0 bool, e1 error, m2 *model.User) {
	defer func() {
		l.logger.Log("method", "GetSpec", "id", id, "b0", b0, "e1", e1, "m2", m2)
	}()
	return l.next.GetSpec(ctx, id)
}
func (l loggingMiddleware) GetAll(ctx context.Context) (b0 bool, e1 error, m2 []model.User) {
	defer func() {
		l.logger.Log("method", "GetAll", "b0", b0, "e1", e1, "m2", m2)
	}()
	return l.next.GetAll(ctx)
}
func (l loggingMiddleware) GetUDF(ctx context.Context, name string) (b0 bool, e1 error, m2 []model.User) {
	defer func() {
		l.logger.Log("method", "GetUDF", "name", name, "b0", b0, "e1", e1, "m2", m2)
	}()
	return l.next.GetUDF(ctx, name)
}
func (l loggingMiddleware) Post(ctx context.Context, user model.User) (b0 bool, e1 error, m2 *model.User) {
	defer func() {
		l.logger.Log("method", "Post", "user", user, "b0", b0, "e1", e1, "m2", m2)
	}()
	return l.next.Post(ctx, user)
}
func (l loggingMiddleware) Patch(ctx context.Context, id string, user model.User) (b0 bool, e1 error, m2 *model.User) {
	defer func() {
		l.logger.Log("method", "Patch", "id", id, "user", user, "b0", b0, "e1", e1, "m2", m2)
	}()
	return l.next.Patch(ctx, id, user)
}
func (l loggingMiddleware) Put(ctx context.Context, id string, user model.User) (b0 bool, e1 error, m2 *model.User) {
	defer func() {
		l.logger.Log("method", "Put", "id", id, "user", user, "b0", b0, "e1", e1, "m2", m2)
	}()
	return l.next.Put(ctx, id, user)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (b0 bool, e1 error, m2 *model.User) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "b0", b0, "e1", e1, "m2", m2)
	}()
	return l.next.Delete(ctx, id)
}
