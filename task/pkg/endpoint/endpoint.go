package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/money-hub/MoneyDodo.service/model"
	service "github.com/money-hub/MoneyDodo.service/task/pkg/service"
)

// GetTasksByIDRequest collects the request parameters for the GetTasksByID method.
type GetTasksByIDRequest struct {
	Id    string `json:"id"`
	State string `json:"state"`
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
		status, errinfo, data := s.GetTasksByID(ctx, req.Id, req.State)
		return GetTasksByIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetTasksByID implements Service. Primarily useful in a client.
func (e Endpoints) GetTasksByID(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Task) {
	request := GetTasksByIDRequest{
		Id:    id,
		State: state,
	}
	response, err := e.GetTasksByIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTasksByIDResponse).Status, response.(GetTasksByIDResponse).Errinfo, response.(GetTasksByIDResponse).Data
}
