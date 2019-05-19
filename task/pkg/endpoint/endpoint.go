package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/money-hub/MoneyDodo.service/model"
	service "github.com/money-hub/MoneyDodo.service/task/pkg/service"
)

// UserGetHisReleasedTasksRequest collects the request parameters for the UserGetHisReleasedTasks method.
type UserGetHisReleasedTasksRequest struct {
	Id string `json:"id"`
}

// UserGetHisReleasedTasksResponse collects the response parameters for the UserGetHisReleasedTasks method.
type UserGetHisReleasedTasksResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeUserGetHisReleasedTasksEndpoint returns an endpoint that invokes UserGetHisReleasedTasks on the service.
func MakeUserGetHisReleasedTasksEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserGetHisReleasedTasksRequest)
		status, errinfo, data := s.UserGetHisReleasedTasks(ctx, req.Id)
		return UserGetHisReleasedTasksResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// UserGetTasksByIDRequest collects the request parameters for the UserGetTasksByID method.
type UserGetTasksByIDRequest struct {
	Id string `json:"id"`
}

// UserGetTasksByIDResponse collects the response parameters for the UserGetTasksByID method.
type UserGetTasksByIDResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeUserGetTasksByIDEndpoint returns an endpoint that invokes UserGetTasksByID on the service.
func MakeUserGetTasksByIDEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserGetTasksByIDRequest)
		status, errinfo, data := s.UserGetTasksByID(ctx, req.Id)
		return UserGetTasksByIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// UserGetHisUnreleasedTasksRequest collects the request parameters for the UserGetHisUnreleasedTasks method.
type UserGetHisUnreleasedTasksRequest struct {
	Id string `json:"id"`
}

// UserGetHisUnreleasedTasksResponse collects the response parameters for the UserGetHisUnreleasedTasks method.
type UserGetHisUnreleasedTasksResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeUserGetHisUnreleasedTasksEndpoint returns an endpoint that invokes UserGetHisUnreleasedTasks on the service.
func MakeUserGetHisUnreleasedTasksEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserGetHisUnreleasedTasksRequest)
		status, errinfo, data := s.UserGetHisUnreleasedTasks(ctx, req.Id)
		return UserGetHisUnreleasedTasksResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// UserGetHisClosedTasksRequest collects the request parameters for the UserGetHisClosedTasks method.
type UserGetHisClosedTasksRequest struct {
	Id string `json:"id"`
}

// UserGetHisClosedTasksResponse collects the response parameters for the UserGetHisClosedTasks method.
type UserGetHisClosedTasksResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeUserGetHisClosedTasksEndpoint returns an endpoint that invokes UserGetHisClosedTasks on the service.
func MakeUserGetHisClosedTasksEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserGetHisClosedTasksRequest)
		status, errinfo, data := s.UserGetHisClosedTasks(ctx, req.Id)
		return UserGetHisClosedTasksResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// AdminGetAllTasksByUserIDRequest collects the request parameters for the AdminGetAllTasksByUserID method.
type AdminGetAllTasksByUserIDRequest struct {
	Id string `json:"id"`
}

// AdminGetAllTasksByUserIDResponse collects the response parameters for the AdminGetAllTasksByUserID method.
type AdminGetAllTasksByUserIDResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeAdminGetAllTasksByUserIDEndpoint returns an endpoint that invokes AdminGetAllTasksByUserID on the service.
func MakeAdminGetAllTasksByUserIDEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AdminGetAllTasksByUserIDRequest)
		status, errinfo, data := s.AdminGetAllTasksByUserID(ctx, req.Id)
		return AdminGetAllTasksByUserIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// AdminGetTasksReleasedByUserIDRequest collects the request parameters for the AdminGetTasksReleasedByUserID method.
type AdminGetTasksReleasedByUserIDRequest struct {
	Id string `json:"id"`
}

// AdminGetTasksReleasedByUserIDResponse collects the response parameters for the AdminGetTasksReleasedByUserID method.
type AdminGetTasksReleasedByUserIDResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeAdminGetTasksReleasedByUserIDEndpoint returns an endpoint that invokes AdminGetTasksReleasedByUserID on the service.
func MakeAdminGetTasksReleasedByUserIDEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AdminGetTasksReleasedByUserIDRequest)
		status, errinfo, data := s.AdminGetTasksReleasedByUserID(ctx, req.Id)
		return AdminGetTasksReleasedByUserIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// AdminGetTasksUnreleasedByUserIDRequest collects the request parameters for the AdminGetTasksUnreleasedByUserID method.
type AdminGetTasksUnreleasedByUserIDRequest struct {
	Id string `json:"id"`
}

// AdminGetTasksUnreleasedByUserIDResponse collects the response parameters for the AdminGetTasksUnreleasedByUserID method.
type AdminGetTasksUnreleasedByUserIDResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeAdminGetTasksUnreleasedByUserIDEndpoint returns an endpoint that invokes AdminGetTasksUnreleasedByUserID on the service.
func MakeAdminGetTasksUnreleasedByUserIDEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AdminGetTasksUnreleasedByUserIDRequest)
		status, errinfo, data := s.AdminGetTasksUnreleasedByUserID(ctx, req.Id)
		return AdminGetTasksUnreleasedByUserIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// AdminGetTasksClosedByUserIDRequest collects the request parameters for the AdminGetTasksClosedByUserID method.
type AdminGetTasksClosedByUserIDRequest struct {
	Id string `json:"id"`
}

// AdminGetTasksClosedByUserIDResponse collects the response parameters for the AdminGetTasksClosedByUserID method.
type AdminGetTasksClosedByUserIDResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Task `json:"data"`
}

// MakeAdminGetTasksClosedByUserIDEndpoint returns an endpoint that invokes AdminGetTasksClosedByUserID on the service.
func MakeAdminGetTasksClosedByUserIDEndpoint(s service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AdminGetTasksClosedByUserIDRequest)
		status, errinfo, data := s.AdminGetTasksClosedByUserID(ctx, req.Id)
		return AdminGetTasksClosedByUserIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// UserGetHisReleasedTasks implements Service. Primarily useful in a client.
func (e Endpoints) UserGetHisReleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := UserGetHisReleasedTasksRequest{Id: id}
	response, err := e.UserGetHisReleasedTasksEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UserGetHisReleasedTasksResponse).Status, response.(UserGetHisReleasedTasksResponse).Errinfo, response.(UserGetHisReleasedTasksResponse).Data
}

// UserGetTasksByID implements Service. Primarily useful in a client.
func (e Endpoints) UserGetTasksByID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := UserGetTasksByIDRequest{Id: id}
	response, err := e.UserGetTasksByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UserGetTasksByIDResponse).Status, response.(UserGetTasksByIDResponse).Errinfo, response.(UserGetTasksByIDResponse).Data
}

// UserGetHisUnreleasedTasks implements Service. Primarily useful in a client.
func (e Endpoints) UserGetHisUnreleasedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := UserGetHisUnreleasedTasksRequest{Id: id}
	response, err := e.UserGetHisUnreleasedTasksEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UserGetHisUnreleasedTasksResponse).Status, response.(UserGetHisUnreleasedTasksResponse).Errinfo, response.(UserGetHisUnreleasedTasksResponse).Data
}

// UserGetHisClosedTasks implements Service. Primarily useful in a client.
func (e Endpoints) UserGetHisClosedTasks(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := UserGetHisClosedTasksRequest{Id: id}
	response, err := e.UserGetHisClosedTasksEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UserGetHisClosedTasksResponse).Status, response.(UserGetHisClosedTasksResponse).Errinfo, response.(UserGetHisClosedTasksResponse).Data
}

// AdminGetAllTasksByUserID implements Service. Primarily useful in a client.
func (e Endpoints) AdminGetAllTasksByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := AdminGetAllTasksByUserIDRequest{Id: id}
	response, err := e.AdminGetAllTasksByUserIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AdminGetAllTasksByUserIDResponse).Status, response.(AdminGetAllTasksByUserIDResponse).Errinfo, response.(AdminGetAllTasksByUserIDResponse).Data
}

// AdminGetTasksReleasedByUserID implements Service. Primarily useful in a client.
func (e Endpoints) AdminGetTasksReleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := AdminGetTasksReleasedByUserIDRequest{Id: id}
	response, err := e.AdminGetTasksReleasedByUserIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AdminGetTasksReleasedByUserIDResponse).Status, response.(AdminGetTasksReleasedByUserIDResponse).Errinfo, response.(AdminGetTasksReleasedByUserIDResponse).Data
}

// AdminGetTasksUnreleasedByUserID implements Service. Primarily useful in a client.
func (e Endpoints) AdminGetTasksUnreleasedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := AdminGetTasksUnreleasedByUserIDRequest{Id: id}
	response, err := e.AdminGetTasksUnreleasedByUserIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AdminGetTasksUnreleasedByUserIDResponse).Status, response.(AdminGetTasksUnreleasedByUserIDResponse).Errinfo, response.(AdminGetTasksUnreleasedByUserIDResponse).Data
}

// AdminGetTasksClosedByUserID implements Service. Primarily useful in a client.
func (e Endpoints) AdminGetTasksClosedByUserID(ctx context.Context, id string) (status bool, errinfo string, data []model.Task) {
	request := AdminGetTasksClosedByUserIDRequest{Id: id}
	response, err := e.AdminGetTasksClosedByUserIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AdminGetTasksClosedByUserIDResponse).Status, response.(AdminGetTasksClosedByUserIDResponse).Errinfo, response.(AdminGetTasksClosedByUserIDResponse).Data
}
