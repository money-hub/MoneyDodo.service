package service

import (
	"context"
	"log"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// DealService describes the service.
type DealService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetUserDealByState(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Deal)
	GetDealByDID(ctx context.Context, id string) (status bool, errinfo string, data model.Deal)
	GetDealByState(ctx context.Context, state string) (status bool, errinfo string, data []model.Deal)
	PostAcceptDeal(ctx context.Context, deal model.Deal) (status bool, errinfo string, data model.Deal)
	PutDealState(ctx context.Context, deal model.Deal) (status bool, errinfo string, data *model.Deal)
}

type basicDealService struct {
	*db.DBService
}

func (b *basicDealService) PutDealState(ctx context.Context, deal model.Deal) (status bool, errinfo string, data *model.Deal) {
	var err error
	d := &model.Deal{
		Id: deal.Id,
	}
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}
	if _, err = sess.Get(d); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	if ctx.Value("id").(string) == deal.Recipient {
		if d.State == model.DealStateUnderway && deal.State == model.DealStateRecConfirm {
			if _, err = sess.Where("id=?", deal.Id).Cols("state").Update(deal); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}
	} else if ctx.Value("id").(string) == deal.Publisher {
		if d.State == model.DealStateUnderway && deal.State == model.DealStateRecConfirm {
			if _, err = sess.Where("id=?", deal.Id).Cols("state").Update(deal); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			// 更新两个用户的余额
			publisher := &model.User{
				Id: d.Publisher,
			}
			if _, err = sess.Get(publisher); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			recipient := &model.User{
				Id: d.Recipient,
			}
			if _, err = sess.Get(recipient); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			publisher.Balance = publisher.Balance + d.Reward
			recipient.Balance = recipient.Balance + d.Reward
			if _, err = sess.Where("id=?", publisher.Id).Cols("balance").Update(publisher); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
			if _, err = sess.Where("id=?", recipient.Id).Cols("balance").Update(recipient); err != nil {
				sess.Rollback()
				return false, err.Error(), nil
			}
		}
	} else {
		sess.Rollback()
		return false, err.Error(), nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	d.State = deal.State
	return true, "", d
}

func (b *basicDealService) GetUserDealByState(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Deal) {
	// TODO implement the business logic of GetUserDealByState
	deal := model.Deal{}
	role := ctx.Value("role")
	userID := ctx.Value("id").(string)
	var rows *xorm.Rows
	var err error
	if role == 0 {
		if state == "" {
			rows, err = b.Engine().Where("publisher = ? or recipient = ?", id, id).Rows(deal)
		} else {
			rows, err = b.Engine().Where("(publisher = ? or recipient = ?) and state = ?", id, id, state).Rows(deal)
		}
		if err == nil {
			for rows.Next() {
				err1 := rows.Scan(&deal)
				if err1 != nil {
					return false, err1.Error(), data
				}
				data = append(data, deal)
			}
			return true, "", data
		}
		return false, err.Error(), data
	} else {
		if userID == id {
			if state == "" {
				rows, err = b.Engine().Where("publisher = ? or recipient = ?", id, id).Rows(deal)
			} else {
				rows, err = b.Engine().Where("(publisher = ? or recipient = ?) and state = ?", id, id, state).Rows(deal)
			}
			if err == nil {
				for rows.Next() {
					err1 := rows.Scan(&deal)
					if err1 != nil {
						return false, err1.Error(), data
					}
					data = append(data, deal)
				}
				return true, "", data
			}
			return false, err.Error(), data
		}
	}
	return false, "Permission denied", data
}
func (b *basicDealService) GetDealByDID(ctx context.Context, id string) (status bool, errinfo string, data model.Deal) {
	// TODO implement the business logic of GetDealByDID
	deal := model.Deal{}
	role := ctx.Value("role")
	userID := ctx.Value("id").(string)
	if role == 0 {
		status, err := b.Engine().Where("id = ?", id).Get(&deal)
		if status == false || err != nil {
			return false, "Get Failed", data
		}
		return true, "", deal
	} else {
		status, err := b.Engine().Where("(publisher = ? or recipient = ?) and id = ?", userID, userID, id).Get(&deal)
		if status == false {
			return false, "Get Failed", data
		} else if err == nil {
			return true, "", deal
		}
	}
	return false, "Permission denied", data
}
func (b *basicDealService) GetDealByState(ctx context.Context, state string) (status bool, errinfo string, data []model.Deal) {
	// TODO implement the business logic of GetDealByState
	deal := model.Deal{}
	role := ctx.Value("role")
	var rows *xorm.Rows
	var err error
	if role == 0 {
		if state != "" {
			rows, err = b.Engine().Where("state = ?", state).Rows(deal)
		} else {
			rows, err = b.Engine().Rows(deal)
		}
		if err == nil {
			for rows.Next() {
				err1 := rows.Scan(&deal)
				if err1 != nil {
					return false, err1.Error(), data
				}
				data = append(data, deal)
			}
			return true, "", data
		}
		return false, err.Error(), data
	}
	return false, "Permission denied", data
}
func (b *basicDealService) PostAcceptDeal(ctx context.Context, deal model.Deal) (status bool, errinfo string, data model.Deal) {
	// TODO implement the business logic of PostAcceptDeal
	role := ctx.Value("role")
	if role == 1 || role == 2 {
		task := model.Task{
			Id: deal.TaskId,
		}
		dealTemp := model.Deal{}
		//task info
		//1. task's id is right
		//2. task hasn't been accepted
		//3. task's publisher is right
		status, err := b.Engine().Where("taskId = ?", deal.TaskId).Get(&dealTemp)
		if status != false {
			return false, "Deal Error", data
		}
		status, err = b.Engine().Where("id = ?", deal.TaskId).Get(&task)
		if status == false || err != nil {
			return false, "Task Error", data
		}
		if task.Publisher != deal.Publisher {
			return false, "Publisher Error", data
		}
		//deal info
		//1. since time
		//2. deal state
		deal.Since = time.Now()
		deal.State = "underway"
		_, err = b.Engine().Insert(deal)
		if err == nil {
			return true, "", deal
		}
		return false, err.Error(), data
	}
	return false, "Permission denied", data
}

// NewBasicDealService returns a naive, stateless implementation of DealService.
func NewBasicDealService() DealService {
	basicDealService := &basicDealService{
		&db.DBService{},
	}
	err := basicDealService.Bind("conf/conf.moneydodo.yml")
	if err != nil {
		log.Printf("The UserService failed to bind with mysql")
	}
	return basicDealService
}

// New returns a DealService with all of the expected middleware wired in.
func New(middleware []Middleware) DealService {
	var svc DealService = NewBasicDealService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
