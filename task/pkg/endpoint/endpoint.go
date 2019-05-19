package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/money-hub/MoneyDodo.service/model"
	service "github.com/money-hub/MoneyDodo.service/task/pkg/service"
)

// GetHisReleasedTasksRequest collects the request parameters for the GetHisReleasedTasks method.
type GetHisReleasedTasksRequest struct {
	Id string `json:"id"`
}

// GetHisReleasedTasksResponse collects the response parameters for the GetHisReleasedTasks method.
type GetHisReleasedTasksResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeGetHisReleasedTasksEndpoint returns an endpoint that invokes GetHisReleasedTasks on the service.
func MakeGetHisReleasedTasksEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetHisReleasedTasksRequest)
		status, errinfo, data := s.GetHisReleasedTasks(ctx, req.Id)
		return GetHisReleasedTasksResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetTasksByIDRequest collects the request parameters for the GetTasksByID method.
type GetTasksByIDRequest struct {
	Id string `json:"id"`
}

// GetTasksByIDResponse collects the response parameters for the GetTasksByID method.
type GetTasksByIDResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeGetTasksByIDEndpoint returns an endpoint that invokes GetTasksByID on the service.
func MakeGetTasksByIDEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTasksByIDRequest)
		status, errinfo, data := s.GetTasksByID(ctx, req.Id)
		return GetTasksByIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetHisUnreleasedTasksRequest collects the request parameters for the GetHisUnreleasedTasks method.
type GetHisUnreleasedTasksRequest struct {
	Id string `json:"id"`
}

// GetHisUnreleasedTasksResponse collects the response parameters for the GetHisUnreleasedTasks method.
type GetHisUnreleasedTasksResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeGetHisUnreleasedTasksEndpoint returns an endpoint that invokes GetHisUnreleasedTasks on the service.
func MakeGetHisUnreleasedTasksEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetHisUnreleasedTasksRequest)
		status, errinfo, data := s.GetHisUnreleasedTasks(ctx, req.Id)
		return GetHisUnreleasedTasksResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetHisClosedTasksRequest collects the request parameters for the GetHisClosedTasks method.
type GetHisClosedTasksRequest struct {
	Id string `json:"id"`
}

// GetHisClosedTasksResponse collects the response parameters for the GetHisClosedTasks method.
type GetHisClosedTasksResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeGetHisClosedTasksEndpoint returns an endpoint that invokes GetHisClosedTasks on the service.
func MakeGetHisClosedTasksEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetHisClosedTasksRequest)
		status, errinfo, data := s.GetHisClosedTasks(ctx, req.Id)
		return GetHisClosedTasksResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetHisReleasedTasks implements Service. Primarily useful in a client.
func (e Endpoints) GetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := GetHisReleasedTasksRequest{Id: id}
	response, err := e.GetHisReleasedTasksEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetHisReleasedTasksResponse).Status, response.(GetHisReleasedTasksResponse).Errinfo, response.(GetHisReleasedTasksResponse).Data
}

// GetTasksByID implements Service. Primarily useful in a client.
func (e Endpoints) GetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := GetTasksByIDRequest{Id: id}
	response, err := e.GetTasksByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTasksByIDResponse).Status, response.(GetTasksByIDResponse).Errinfo, response.(GetTasksByIDResponse).Data
}

// GetHisUnreleasedTasks implements Service. Primarily useful in a client.
func (e Endpoints) GetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := GetHisUnreleasedTasksRequest{Id: id}
	response, err := e.GetHisUnreleasedTasksEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetHisUnreleasedTasksResponse).Status, response.(GetHisUnreleasedTasksResponse).Errinfo, response.(GetHisUnreleasedTasksResponse).Data
}

// GetHisClosedTasks implements Service. Primarily useful in a client.
func (e Endpoints) GetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := GetHisClosedTasksRequest{Id: id}
	response, err := e.GetHisClosedTasksEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetHisClosedTasksResponse).Status, response.(GetHisClosedTasksResponse).Errinfo, response.(GetHisClosedTasksResponse).Data
}
