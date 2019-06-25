package charge

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
)

// 此函数并没有使用，函数已在go-kit包中实现
func DelayingLimiterMiddleware(limit ratelimit.Waiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if err := limit.Wait(ctx); err != nil {
				return nil, err
			}
			return next(ctx, request)
		}
	}
}
