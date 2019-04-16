package service

import (
	"context"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// UserService describes the service.
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetSpec(ctx context.Context, id string) (bool, error, interface{})
	GetAll(ctx context.Context) (bool, error, interface{})
	GetUDF(ctx context.Context, name string) (bool, error, interface{})
	Post(ctx context.Context, user model.User) (bool, error, interface{})
	Patch(ctx context.Context, user model.User) (bool, error, interface{})
	Put(ctx context.Context, user model.User) (bool, error, interface{})
	Delete(ctx context.Context, id string) (bool, error, interface{})
}

type basicUserService struct {
	*db.DBService
}

func (b *basicUserService) GetSpec(ctx context.Context, id string) (bool, error, interface{}) {
	// TODO implement the business logic of GetSpec
	user := &model.User{
		Id: id,
	}
	ok, err := b.Engine().Get(user)
	log.Println(ok, err)
	return ok, err, user
}
func (b *basicUserService) GetAll(ctx context.Context) (b0 bool, e1 error, i2 interface{}) {
	// TODO implement the business logic of GetAll
	return b0, e1, i2
}
func (b *basicUserService) GetUDF(ctx context.Context, name string) (b0 bool, e1 error, i2 interface{}) {
	// TODO implement the business logic of GetUDF
	return b0, e1, i2
}
func (b *basicUserService) Post(ctx context.Context, user model.User) (b0 bool, e1 error, i2 interface{}) {
	// TODO implement the business logic of Post
	return b0, e1, i2
}
func (b *basicUserService) Patch(ctx context.Context, user model.User) (b0 bool, e1 error, i2 interface{}) {
	// TODO implement the business logic of Patch
	return b0, e1, i2
}
func (b *basicUserService) Put(ctx context.Context, user model.User) (b0 bool, e1 error, i2 interface{}) {
	// TODO implement the business logic of Put
	return b0, e1, i2
}
func (b *basicUserService) Delete(ctx context.Context, id string) (b0 bool, e1 error, i2 interface{}) {
	// TODO implement the business logic of Delete
	return b0, e1, i2
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	basicUserSvc := &basicUserService{
		&db.DBService{},
	}
	err := basicUserSvc.Bind("conf/conf.user.yml")
	if err != nil {
		log.Printf("The UserService failed to bind with mysql")
	}
	return basicUserSvc
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
