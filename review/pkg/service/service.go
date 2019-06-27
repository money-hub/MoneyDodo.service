package service

import (
	"context"
	"fmt"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// ReviewService describes the service.
type ReviewService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	// 商家提交审核
	PostReview(ctx context.Context, review model.Review) (status bool, errinfo string, data model.Review)

	// 查看指定id审核结果
	GetReview(ctx context.Context, rid string) (status bool, errinfo string, data model.Review)

	// 查询所有审核结果
	GetReviews(ctx context.Context) (status bool, errinfo string, data []model.Review)

	// 修改审核状态
	PutReview(ctx context.Context, rid string, review model.Review) (status bool, errinfo string, data model.Review)
}

// 商家提交审核
func (b *basicReviewService) PostReview(ctx context.Context, review model.Review) (status bool, errinfo string, data model.Review) {
	// TODO implement the business logic of PostReview
	if ctx.Value("role").(int) == 2 {
		_, err := b.Engine().Insert(review)
		if err != nil {
			fmt.Println(err.Error())
			return false, err.Error(), review
		}
		b.Engine().Desc("id").Limit(1, 0).Get(review)
		return true, "", review
	}
	return false, "This service only for enterprise", review
}

// 查看指定id审核结果
func (b *basicReviewService) GetReview(ctx context.Context, rid string) (status bool, errinfo string, data model.Review) {
	// TODO implement the business logic of GetReview
	review := model.Review{Id: rid}
	if ctx.Value("role").(int) == 2 {
		if has, _ := b.Engine().Get(&review); has == false {
			return false, "No such a review", review
		}
		return true, "", review
	}
	return false, "This service only for enterprise", review
}

// 查询所有审核结果
func (b *basicReviewService) GetReviews(ctx context.Context) (status bool, errinfo string, data []model.Review) {
	// TODO implement the business logic of GetReviews

	reviews := make([]model.Review, 0)
	if ctx.Value("role").(int) == 2 {
		// 商家
		err := b.Engine().Where("name = ?", ctx.Value("id")).Find(&reviews)
		if err != nil {
			log.Println(err)
			return false, err.Error(), nil
		}
		return true, "", reviews

	} else if ctx.Value("role").(int) == 0 {
		// 管理员
		err := b.Engine().Find(&reviews)
		if err != nil {
			log.Println(err)
			return false, err.Error(), nil
		}
		return true, "", reviews
	}
	return false, "This service only for enterprise or admin", reviews
}

// 修改审核状态
func (b *basicReviewService) PutReview(ctx context.Context, rid string, review model.Review) (status bool, errinfo string, data model.Review) {
	// TODO implement the business logic of PutReview
	if ctx.Value("role").(int) == 0 {
		if has, _ := b.Engine().Get(&review); has == false {
			return false, "No such a review", review
		}
		_, err := b.Engine().Where("id=?", rid).Update(review)
		if err != nil {
			fmt.Println(err.Error())
			return false, err.Error(), review
		}
		return true, "", review
	}
	return false, "No permission", review
}

type basicReviewService struct {
	*db.DBService
}

// NewBasicReviewService returns a naive, stateless implementation of ReviewService.
func NewBasicReviewService() ReviewService {
	basicReviewSvc := &basicReviewService{
		&db.DBService{},
	}

	err := basicReviewSvc.Bind("conf/conf.moneydodo.yml")
	if err != nil {
		log.Printf("The ReviewService failed to bind with mysql")
	}
	return basicReviewSvc
}

// New returns a ReviewService with all of the expected middleware wired in.
func New(middleware []Middleware) ReviewService {
	var svc ReviewService = NewBasicReviewService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
