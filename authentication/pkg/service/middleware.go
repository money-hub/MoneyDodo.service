package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthenticationService) AuthenticationService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthenticationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthenticationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthenticationService) AuthenticationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetOpenId(ctx context.Context, code string) (e0 error, s1 string, s2 string) {
	defer func() {
		l.logger.Log("method", "GetOpenId", "code", code, "e0", e0, "s1", s1, "s2", s2)
	}()
	return l.next.GetOpenId(ctx, code)
}
func (l loggingMiddleware) AdminLogin(ctx context.Context, username string, password string) (e0 error, b1 bool, s2 string) {
	defer func() {
		l.logger.Log("method", "AdminLogin", "username", username, "password", password, "e0", e0, "b1", b1, "s2", s2)
	}()
	return l.next.AdminLogin(ctx, username, password)
}
