package personalTasks

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/money-hub/MoneyDodo.service/model"
)

type Endpoints struct {
	GetSpecEndpoint   endpoint.Endpoint
	GetAllEndpoint    endpoint.Endpoint
	PostClaimEndpoint endpoint.Endpoint
	PostEndpoint      endpoint.Endpoint
	PutEndpoint       endpoint.Endpoint
	DeleteEndpoint    endpoint.Endpoint
}

func MakeServerEndpoints(s PTaskService) Endpoints {
	eps := Endpoints{
		GetSpecEndpoint:   MakeGetSpecEndpoint(s),
		GetAllEndpoint:    MakeGetAllEndpoint(s),
		PostClaimEndpoint: MakePostClaimEndpoint(s),
		PostEndpoint:      MakePostEndpoint(s),
		PutEndpoint:       MakePutEndpoint(s),
		DeleteEndpoint:    MakeDeleteEndpoint(s),
	}
	return eps
}

// GetSpecRequest collects the request parameters for the GetSpec method.
type GetSpecRequest struct {
	UserId string `json:"userId"`
	TaskId string `json:"taskId"`
}

// GetSpecResponse collects the response parameters for the GetSpec method.
type GetSpecResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.Task `json:"data"`
}

// MakeGetSpecEndpoint returns an endpoint that invokes GetSpec on the service.
func MakeGetSpecEndpoint(s PTaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSpecRequest)
		status, errinfo, data := s.GetSpec(ctx, req.UserId, req.TaskId)
		return GetSpecResponse{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// GetAllRequest collects the request parameters for the GetAll method.
type GetAllRequest struct {
	UserId string `json:"userId"`
}

// GetAllResponse collects the response parameters for the GetAll method.
type GetAllResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s PTaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRequest)
		status, errinfo, data := s.GetAll(ctx, req.UserId)
		return GetAllResponse{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// PostRequest collects the response parameters for the Post method.
type PostRequest struct {
	UserId string     `json:"userId"`
	Task   model.Task `json:"task"`
}

// PostResponse collects the response parameters for the Post method.
type PostResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.Task `json:"data"`
}

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s PTaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		status, errinfo, data := s.Post(ctx, req.UserId, req.Task)
		return PostResponse{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// PostClaimRequest collects the response parameters for the PostClaim method.
type PostClaimRequest struct {
	UserId string `json:"userId"`
	TaskId string `json:"taskId"`
}

// PostClaimResponse collects the response parameters for the PostClaim method.
type PostClaimResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.Task `json:"data"`
}

// MakePostClaimEndpoint returns an endpoint that invokes PostClaim on the service.
func MakePostClaimEndpoint(s PTaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostClaimRequest)
		status, errinfo, data := s.PostClaim(ctx, req.UserId, req.TaskId)
		return PostClaimResponse{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// DeleteRequest collects the response parameters for the Delete method.
type DeleteRequest struct {
	UserId string `json:"userId"`
	TaskId string `json:"taskId"`
	Detail string `json:"detail"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.Task `json:"data"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s PTaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		status, errinfo, data := s.Delete(ctx, req.UserId, req.TaskId, req.Detail)
		return DeleteResponse{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}

// PutRequest collects the response parameters for the Put method.
type PutRequest struct {
	UserId string     `json:"userId"`
	TaskId string     `json:"taskId"`
	Task   model.Task `json:"task"`
}

// PutResponse collects the response parameters for the Put method.
type PutResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.Task `json:"data"`
}

// MakePutEndpoint returns an endpoint that invokes Put on the service.
func MakePutEndpoint(s PTaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutRequest)
		status, errinfo, data := s.Put(ctx, req.UserId, req.TaskId, req.Task)
		return PutResponse{
			Status:  status,
			Errinfo: errinfo,
			Data:    data,
		}, nil
	}
}
