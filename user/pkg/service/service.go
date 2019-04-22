package service

import (
	"context"
	"errors"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// UserService describes the service.
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetSpec(ctx context.Context, id string) (bool, error, *model.User)
	GetAll(ctx context.Context) (bool, error, []model.User)
	GetUDF(ctx context.Context, name string) (bool, error, []model.User)
	Post(ctx context.Context, user model.User) (bool, error, *model.User)
	Patch(ctx context.Context, id string, user model.User) (bool, error, *model.User)
	Put(ctx context.Context, id string, user model.User) (bool, error, *model.User)
	Delete(ctx context.Context, id string) (bool, error, *model.User)
}

type basicUserService struct {
	*db.DBService
}

func (b *basicUserService) GetSpec(ctx context.Context, id string) (bool, error, *model.User) {
	// TODO implement the business logic of GetSpec
	user := &model.User{
		Id: id,
	}
	ok, err := b.Engine().Get(user)
	if ok == false {
		user = nil
	}
	return ok, err, user
}
func (b *basicUserService) GetAll(ctx context.Context) (bool, error, []model.User) {
	// TODO implement the business logic of GetAll
	users := make([]model.User, 0)
	err := b.Engine().Find(&users)
	ok := true
	if err != nil {
		ok = false
	}
	return ok, err, users
}
func (b *basicUserService) GetUDF(ctx context.Context, name string) (bool, error, []model.User) {
	// TODO implement the business logic of GetUDF
	_, err, users := b.GetAll(ctx)
	udfUsers := make([]model.User, 0)
	for _, user := range users {
		if user.Name == name {
			udfUsers = append(udfUsers, user)
		}
	}
	return len(udfUsers) > 0, err, udfUsers
}
func (b *basicUserService) Post(ctx context.Context, user model.User) (bool, error, *model.User) {
	// TODO implement the business logic of Post
	if user.Id == "" {
		return false, errors.New("not a valid userid"), nil
	}
	row, err := b.Engine().Insert(user)
	return row > 0, err, nil
}
func (b *basicUserService) Patch(ctx context.Context, id string, user model.User) (b0 bool, e1 error, i2 *model.User) {
	// TODO implement the business logic of Patch
	return b0, e1, i2
}
func (b *basicUserService) Put(ctx context.Context, id string, user model.User) (bool, error, *model.User) {
	// TODO implement the business logic of Put
	row, err := b.Engine().Where("id = ?", id).AllCols().Update(user)
	return row > 0, err, nil
}
func (b *basicUserService) Delete(ctx context.Context, id string) (bool, error, *model.User) {
	// TODO implement the business logic of Delete
	user := model.User{
		Id: id,
	}
	row, err := b.Engine().Delete(user)
	return row > 0, err, nil
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
