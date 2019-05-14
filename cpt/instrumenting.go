package cpt

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/money-hub/MoneyDodo.service/model"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           CptService
}

func (i *InstrumentingMiddleware) GetSpec(ctx context.Context, taskId string) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpec", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.GetSpec(ctx, taskId)
	return
}

func (i *InstrumentingMiddleware) GetAll(ctx context.Context, kind, state string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Task) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAll", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.GetAll(ctx, kind, state, page, offset, limit, orderby)
	return
}

func (i *InstrumentingMiddleware) Post(ctx context.Context, kind string, task interface{}) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Post", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.Post(ctx, kind, task)
	return
}

func (i *InstrumentingMiddleware) Put(ctx context.Context, taskId string, task interface{}) (status bool, errinfo string, data interface{}) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Put", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.Put(ctx, taskId, task)
	return
}

func (i *InstrumentingMiddleware) Delete(ctx context.Context, taskId string, state string) (status bool, errinfo string, data *model.Task) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Delete", "error", errinfo}
		i.RequestCount.With(lvs...).Add(1)
		i.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	status, errinfo, data = i.Next.Delete(ctx, taskId, state)
	return
}
