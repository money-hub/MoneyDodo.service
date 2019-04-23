package service

import "context"

// AuthenticationService describes the service.
type AuthenticationService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetOpenId(ctx context.Context, code string) (error, string, string)

	AdminLogin(ctx context.Context, username string, password string) (error, bool, string)
}

type basicAuthenticationService struct{}

func (b *basicAuthenticationService) GetOpenId(ctx context.Context, code string) (e0 error, s1 string, s2 string) {
	// TODO implement the business logic of GetOpenId
	return e0, s1, s2
}
func (b *basicAuthenticationService) AdminLogin(ctx context.Context, username string, password string) (e0 error, b1 bool, s2 string) {
	// TODO implement the business logic of AdminLogin
	return e0, b1, s2
}

// NewBasicAuthenticationService returns a naive, stateless implementation of AuthenticationService.
func NewBasicAuthenticationService() AuthenticationService {
	return &basicAuthenticationService{}
}

// New returns a AuthenticationService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthenticationService {
	var svc AuthenticationService = NewBasicAuthenticationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
