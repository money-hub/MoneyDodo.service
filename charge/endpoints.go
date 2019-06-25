package charge

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetSpecEndpoint      endpoint.Endpoint
	GetAllEndpoint       endpoint.Endpoint
	GetAllOfUserEndpoint endpoint.Endpoint
	PostEndpoint         endpoint.Endpoint
	DeleteEndpoint       endpoint.Endpoint
}

func MakeServerEndpoints(s ChargeService) Endpoints {
	eps := Endpoints{
		GetSpecEndpoint:      MakeGetSpecEndpoint(s),
		GetAllEndpoint:       MakeGetAllEndpoint(s),
		GetAllOfUserEndpoint: MakeGetAllOfUserEndpoint(s),
		PostEndpoint:         MakePostEndpoint(s),
		DeleteEndpoint:       MakeDeleteEndpoint(s),
	}
	return eps
}

// Request collects the request parameters for the All method.
type Request struct {
	ChargeId string      `json:"chargeId"`
	Charge   interface{} `json:"charge"`
	UserId   string      `json:"userId"`
	Page     int         `json:"page"`
	Offset   int         `json:"offset"`
	Limit    int         `json:"limit"`
	Orderby  string      `json:"orderby"`
}

// Response collects the response parameters for the All method.
type Response struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    interface{} `json:"data"`
}

// MakeGetSpecEndpoint returns an endpoint that invokes GetSpec on the service.
func MakeGetSpecEndpoint(s ChargeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.GetSpec(ctx, req.ChargeId)
		return Response{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// // GetAllRequest collects the request parameters for the GetAll method.
// type GetAllRequest struct {
// 	UserId string `json:"userId"`
// }

// // GetAllResponse collects the response parameters for the GetAll method.
// type GetAllResponse struct {
// 	Status  bool         `json:"status"`
// 	Errinfo string       `json:"errinfo"`
// 	Data    []model.Task `json:"data"`
// }

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s ChargeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.GetAll(ctx, req.Page, req.Offset, req.Limit, req.Orderby)
		return Response{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

func MakeGetAllOfUserEndpoint(s ChargeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.GetAllOfUser(ctx, req.UserId, req.Page, req.Offset, req.Limit, req.Orderby)
		return Response{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// // PostRequest collects the response parameters for the Post method.
// type PostRequest struct {
// 	UserId string     `json:"userId"`
// 	Task   model.Task `json:"task"`
// }

// // PostResponse collects the response parameters for the Post method.
// type PostResponse struct {
// 	Status  bool        `json:"status"`
// 	Errinfo string      `json:"errinfo"`
// 	Data    *model.Task `json:"data"`
// }

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s ChargeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.Post(ctx, req.Charge)
		return Response{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// // DeleteRequest collects the response parameters for the Delete method.
// type DeleteRequest struct {
// 	UserId string `json:"userId"`
// 	TaskId string `json:"taskId"`
// 	Status string `json:"status"`
// }

// // DeleteResponse collects the response parameters for the Delete method.
// type DeleteResponse struct {
// 	Status  bool        `json:"status"`
// 	Errinfo string      `json:"errinfo"`
// 	Data    *model.Task `json:"data"`
// }

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s ChargeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.Delete(ctx, req.ChargeId)
		return Response{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}
