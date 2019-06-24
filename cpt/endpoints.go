package cpt

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetSpecEndpoint endpoint.Endpoint
	GetAllEndpoint  endpoint.Endpoint
	PostEndpoint    endpoint.Endpoint
	PutEndpoint     endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(s CptService) Endpoints {
	eps := Endpoints{
		GetSpecEndpoint: MakeGetSpecEndpoint(s),
		GetAllEndpoint:  MakeGetAllEndpoint(s),
		PostEndpoint:    MakePostEndpoint(s),
		PutEndpoint:     MakePutEndpoint(s),
		DeleteEndpoint:  MakeDeleteEndpoint(s),
	}
	return eps
}

// Request collects the request parameters for the All method.
type Request struct {
	TaskId  string      `json:"taskId"`
	Kind    string      `json:"kind"`
	State   string      `json:"state"`
	Task    interface{} `json:"task"`
	Page    int         `json:"page"`
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Orderby string      `json:"orderby"`
}

// Response collects the response parameters for the All method.
type Response struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    interface{} `json:"data"`
}

// MakeGetSpecEndpoint returns an endpoint that invokes GetSpec on the service.
func MakeGetSpecEndpoint(s CptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.GetSpec(ctx, req.TaskId)
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
func MakeGetAllEndpoint(s CptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.GetAll(ctx, req.Kind, req.State, req.Page, req.Offset, req.Limit, req.Orderby)
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
func MakePostEndpoint(s CptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.Post(ctx, req.Kind, req.Task)
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
func MakeDeleteEndpoint(s CptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.Delete(ctx, req.TaskId, req.State)
		return Response{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// // PutRequest collects the response parameters for the Put method.
// type PutRequest struct {
// 	UserId string     `json:"userId"`
// 	TaskId string     `json:"taskId"`
// 	Task   model.Task `json:"task"`
// }

// // PutResponse collects the response parameters for the Put method.
// type PutResponse struct {
// 	Status  bool        `json:"status"`
// 	Errinfo string      `json:"errinfo"`
// 	Data    *model.Task `json:"data"`
// }

// MakePutEndpoint returns an endpoint that invokes Put on the service.
func MakePutEndpoint(s CptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		status, errinfo, data := s.Put(ctx, req.TaskId, req.Task)
		return Response{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}
