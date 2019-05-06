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
	GetSpec(ctx context.Context, id string) (status bool, errinfo string, data *model.User)
	GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.User)
	GetUDF(ctx context.Context, name string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.User)
	Post(ctx context.Context, user model.User) (status bool, errinfo string, data *model.User)
	Patch(ctx context.Context, id string, user model.User) (status bool, errinfo string, data *model.User)
	Put(ctx context.Context, id string, user model.User) (status bool, errinfo string, data *model.User)
	Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User)
}

type basicUserService struct {
	*db.DBService
}

func (b *basicUserService) GetSpec(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	// TODO implement the business logic of GetSpec
	data = &model.User{
		Id: id,
	}
	status, err := b.Engine().Get(data)
	if status == false {
		data = nil
	}
	if err == nil {
		errinfo = ""
	} else {
		errinfo = err.Error()
	}
	return
}
func (b *basicUserService) GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.User) {
	// TODO implement the business logic of GetAll
	data = make([]model.User, 0)
	err := b.Engine().Find(&data)

	if orderby == "-id" {
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}

	if limit > 0 {
		res := make([]model.User, 0)
		for i, user := range data {
			if i >= offset && i < offset+page*limit {
				res = append(res, user)
			}
		}

		status = len(res) > 0
		if err != nil {
			errinfo = err.Error()
		}
		return status, errinfo, res
	}

	status = len(data) > 0
	if err != nil {
		errinfo = err.Error()
	}
	return
}
func (b *basicUserService) GetUDF(ctx context.Context, name string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.User) {
	// TODO implement the business logic of GetUDF
	_, err, users := b.GetAll(ctx, page, offset, limit, orderby)
	udfUsers := make([]model.User, 0)
	for _, user := range users {
		if user.Name == name {
			udfUsers = append(udfUsers, user)
		}
	}
	return len(udfUsers) > 0, err, udfUsers
}
func (b *basicUserService) Post(ctx context.Context, user model.User) (status bool, errinfo string, data *model.User) {
	// TODO implement the business logic of Post
	if user.Id == "" {
		return false, "not a valid userid", nil
	}
	row, err := b.Engine().Insert(user)
	if err == nil {
		errinfo = ""
	} else {
		errinfo = err.Error()
	}
	return row > 0, errinfo, nil
}
func (b *basicUserService) Patch(ctx context.Context, id string, user model.User) (b0 bool, e1 string, i2 *model.User) {
	// TODO implement the business logic of Patch
	return b0, e1, i2
}
func (b *basicUserService) Put(ctx context.Context, id string, user model.User) (status bool, errinfo string, data *model.User) {
	// TODO implement the business logic of Put
	if id != user.Id {
		return false, "The user.Id can not be modified.", nil
	}
	row, err := b.Engine().Where("id = ?", id).AllCols().Update(user)
	if err == nil {
		errinfo = ""
	} else {
		errinfo = err.Error()
	}
	return row > 0, errinfo, nil
}
func (b *basicUserService) Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	// TODO implement the business logic of Delete
	user := model.User{
		Id: id,
	}
	row, err := b.Engine().Delete(user)
	if err == nil {
		errinfo = ""
	} else {
		errinfo = err.Error()
	}
	return row > 0, errinfo, nil
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	basicUserSvc := &basicUserService{
		&db.DBService{},
	}
	err := basicUserSvc.Bind("conf/conf.lyh.yml")
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
