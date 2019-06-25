package charge

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/money-hub/MoneyDodo.service/model"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           ChargeService
}

func (i *InstrumentingMiddleware) GetSpec(ctx context.Context, chargeId string) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpec", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.GetSpec(ctx, chargeId)
	return
}

func (i *InstrumentingMiddleware) GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAll", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.GetAll(ctx, page, offset, limit, orderby)
	return
}

func (i *InstrumentingMiddleware) GetAllOfUser(ctx context.Context, userId string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAll", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.GetAllOfUser(ctx, userId, page, offset, limit, orderby)
	return
}

func (i *InstrumentingMiddleware) Post(ctx context.Context, charge interface{}) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Post", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.Post(ctx, charge)
	return
}

func (i *InstrumentingMiddleware) Delete(ctx context.Context, chargeId string) (status bool, errinfo string, data *model.Charge) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Delete", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.Delete(ctx, chargeId)
	return
}
