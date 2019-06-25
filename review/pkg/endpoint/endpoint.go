package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	model "github.com/money-hub/MoneyDodo.service/model"
	service "github.com/money-hub/MoneyDodo.service/review/pkg/service"
)

// PostReviewRequest collects the request parameters for the PostReview method.
type PostReviewRequest struct {
	Review model.Review `json:"review"`
}

// PostReviewResponse collects the response parameters for the PostReview method.
type PostReviewResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    model.Review `json:"data"`
}

// MakePostReviewEndpoint returns an endpoint that invokes PostReview on the service.
func MakePostReviewEndpoint(s service.ReviewService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostReviewRequest)
		status, errinfo, data := s.PostReview(ctx, req.Review)
		return PostReviewResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetReviewRequest collects the request parameters for the GetReview method.
type GetReviewRequest struct {
	Rid string `json:"rid"`
}

// GetReviewResponse collects the response parameters for the GetReview method.
type GetReviewResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    model.Review `json:"data"`
}

// MakeGetReviewEndpoint returns an endpoint that invokes GetReview on the service.
func MakeGetReviewEndpoint(s service.ReviewService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetReviewRequest)
		status, errinfo, data := s.GetReview(ctx, req.Rid)
		return GetReviewResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetReviewsRequest collects the request parameters for the GetReviews method.
type GetReviewsRequest struct{}

// GetReviewsResponse collects the response parameters for the GetReviews method.
type GetReviewsResponse struct {
	Status  bool           `json:"status"`
	Errinfo string         `json:"errinfo"`
	Data    []model.Review `json:"data"`
}

// MakeGetReviewsEndpoint returns an endpoint that invokes GetReviews on the service.
func MakeGetReviewsEndpoint(s service.ReviewService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status, errinfo, data := s.GetReviews(ctx)
		return GetReviewsResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PutReviewRequest collects the request parameters for the PutReview method.
type PutReviewRequest struct {
	Rid    string       `json:"rid"`
	Review model.Review `json:"review"`
}

// PutReviewResponse collects the response parameters for the PutReview method.
type PutReviewResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    model.Review `json:"data"`
}

// MakePutReviewEndpoint returns an endpoint that invokes PutReview on the service.
func MakePutReviewEndpoint(s service.ReviewService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutReviewRequest)
		status, errinfo, data := s.PutReview(ctx, req.Rid, req.Review)
		return PutReviewResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostReview implements Service. Primarily useful in a client.
func (e Endpoints) PostReview(ctx context.Context, review model.Review) (status bool, errinfo string, data model.Review) {
	request := PostReviewRequest{Review: review}
	response, err := e.PostReviewEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostReviewResponse).Status, response.(PostReviewResponse).Errinfo, response.(PostReviewResponse).Data
}

// GetReview implements Service. Primarily useful in a client.
func (e Endpoints) GetReview(ctx context.Context, rid string) (status bool, errinfo string, data model.Review) {
	request := GetReviewRequest{Rid: rid}
	response, err := e.GetReviewEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetReviewResponse).Status, response.(GetReviewResponse).Errinfo, response.(GetReviewResponse).Data
}

// GetReviews implements Service. Primarily useful in a client.
func (e Endpoints) GetReviews(ctx context.Context) (status bool, errinfo string, data []model.Review) {
	request := GetReviewsRequest{}
	response, err := e.GetReviewsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetReviewsResponse).Status, response.(GetReviewsResponse).Errinfo, response.(GetReviewsResponse).Data
}

// PutReview implements Service. Primarily useful in a client.
func (e Endpoints) PutReview(ctx context.Context, rid string, review model.Review) (status bool, errinfo string, data model.Review) {
	request := PutReviewRequest{
		Review: review,
		Rid:    rid,
	}
	response, err := e.PutReviewEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PutReviewResponse).Status, response.(PutReviewResponse).Errinfo, response.(PutReviewResponse).Data
}
