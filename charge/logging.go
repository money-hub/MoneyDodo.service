package charge

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/money-hub/MoneyDodo.service/model"
)

// log middleware
type LoggingMiddleware struct {
	Logger log.Logger
	Next   ChargeService
}

func (l *LoggingMiddleware) GetSpec(ctx context.Context, chargeId string) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetSpec",
			"input", fmt.Sprintf("%#v", Request{ChargeId: chargeId}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetSpec(ctx, chargeId)
	return
}

func (l *LoggingMiddleware) GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetAll",
			"input", fmt.Sprintf("%#v", Request{Page: page, Offset: offset, Limit: limit, Orderby: orderby}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetAll(ctx, page, offset, limit, orderby)
	return
}

func (l *LoggingMiddleware) GetAllOfUser(ctx context.Context, userId string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "GetAll",
			"input", fmt.Sprintf("%#v", Request{UserId: userId, Page: page, Offset: offset, Limit: limit, Orderby: orderby}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.GetAllOfUser(ctx, userId, page, offset, limit, orderby)
	return
}

func (l *LoggingMiddleware) Post(ctx context.Context, charge interface{}) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Post",
			"input", fmt.Sprintf("%#v", Request{Charge: charge}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Post(ctx, charge)
	return
}

func (l *LoggingMiddleware) Delete(ctx context.Context, chargeId string) (status bool, errinfo string, data *model.Charge) {
	defer func(begin time.Time) {
		l.Logger.Log(
			"method", "Delete",
			"input", fmt.Sprintf("%#v", Request{ChargeId: chargeId}),
			"output", fmt.Sprintf("%#v", Response{status, errinfo, data}),
			"took", time.Since(begin),
		)
	}(time.Now())
	status, errinfo, data = l.Next.Delete(ctx, chargeId)
	return
}
