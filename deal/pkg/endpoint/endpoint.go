package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/money-hub/MoneyDodo.service/deal/pkg/service"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// GetUserDealByStateRequest collects the request parameters for the GetUserDealByState method.
type GetUserDealByStateRequest struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

// GetUserDealByStateResponse collects the response parameters for the GetUserDealByState method.
type GetUserDealByStateResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Deal `json:"data"`
}

// MakeGetUserDealByStateEndpoint returns an endpoint that invokes GetUserDealByState on the service.
func MakeGetUserDealByStateEndpoint(s service.DealService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserDealByStateRequest)
		status, errinfo, data := s.GetUserDealByState(ctx, req.Id, req.State)
		return GetUserDealByStateResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetDealByDIDRequest collects the request parameters for the GetDealByDID method.
type GetDealByDIDRequest struct {
	Id string `json:"id"`
}

// GetDealByDIDResponse collects the response parameters for the GetDealByDID method.
type GetDealByDIDResponse struct {
	Status  bool       `json:"status"`
	Errinfo string     `json:"errinfo"`
	Data    model.Deal `json:"data"`
}

// MakeGetDealByDIDEndpoint returns an endpoint that invokes GetDealByDID on the service.
func MakeGetDealByDIDEndpoint(s service.DealService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDealByDIDRequest)
		status, errinfo, data := s.GetDealByDID(ctx, req.Id)
		return GetDealByDIDResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetDealByStateRequest collects the request parameters for the GetDealByState method.
type GetDealByStateRequest struct {
	State string `json:"state"`
}

// GetDealByStateResponse collects the response parameters for the GetDealByState method.
type GetDealByStateResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.Deal `json:"data"`
}

// MakeGetDealByStateEndpoint returns an endpoint that invokes GetDealByState on the service.
func MakeGetDealByStateEndpoint(s service.DealService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDealByStateRequest)
		status, errinfo, data := s.GetDealByState(ctx, req.State)
		return GetDealByStateResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostAcceptDealRequest collects the request parameters for the PostAcceptDeal method.
type PostAcceptDealRequest struct {
	Deal model.Deal `json:"deal"`
}

// PostAcceptDealResponse collects the response parameters for the PostAcceptDeal method.
type PostAcceptDealResponse struct {
	Status  bool       `json:"status"`
	Errinfo string     `json:"errinfo"`
	Data    model.Deal `json:"data"`
}

// MakePostAcceptDealEndpoint returns an endpoint that invokes PostAcceptDeal on the service.
func MakePostAcceptDealEndpoint(s service.DealService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostAcceptDealRequest)
		status, errinfo, data := s.PostAcceptDeal(ctx, req.Deal)
		return PostAcceptDealResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PutDealStateRequest collects the request parameters for the PutDealState method.
type PutDealStateRequest struct {
	Deal model.Deal `json:"deal"`
}

// PutDealStateResponse collects the response parameters for the PutDealState method.
type PutDealStateResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.Deal `json:"data"`
}

// MakePutDealStateEndpoint returns an endpoint that invokes PutDealState on the service.
func MakePutDealStateEndpoint(s service.DealService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutDealStateRequest)
		status, errinfo, data := s.PutDealState(ctx, req.Deal)
		return PutDealStateResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetUserDealByState implements Service. Primarily useful in a client.
func (e Endpoints) GetUserDealByState(ctx context.Context, id string, state string) (status bool, errinfo string, data []model.Deal) {
	request := GetUserDealByStateRequest{
		Id:    id,
		State: state,
	}
	response, err := e.GetUserDealByStateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserDealByStateResponse).Status, response.(GetUserDealByStateResponse).Errinfo, response.(GetUserDealByStateResponse).Data
}

// GetDealByDID implements Service. Primarily useful in a client.
func (e Endpoints) GetDealByDID(ctx context.Context, id string) (status bool, errinfo string, data model.Deal) {
	request := GetDealByDIDRequest{Id: id}
	response, err := e.GetDealByDIDEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetDealByDIDResponse).Status, response.(GetDealByDIDResponse).Errinfo, response.(GetDealByDIDResponse).Data
}

// GetDealByState implements Service. Primarily useful in a client.
func (e Endpoints) GetDealByState(ctx context.Context, state string) (status bool, errinfo string, data []model.Deal) {
	request := GetDealByStateRequest{State: state}
	response, err := e.GetDealByStateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetDealByStateResponse).Status, response.(GetDealByStateResponse).Errinfo, response.(GetDealByStateResponse).Data
}

// PostAcceptDeal implements Service. Primarily useful in a client.
func (e Endpoints) PostAcceptDeal(ctx context.Context, deal model.Deal) (status bool, errinfo string, data model.Deal) {
	request := PostAcceptDealRequest{Deal: deal}
	response, err := e.PostAcceptDealEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostAcceptDealResponse).Status, response.(PostAcceptDealResponse).Errinfo, response.(PostAcceptDealResponse).Data
}

// PutDealState implements Service. Primarily useful in a client.
func (e Endpoints) PutDealState(ctx context.Context, deal model.Deal) (status bool, errinfo string, data *model.Deal) {
	request := PutDealStateRequest{Deal: deal}
	response, err := e.PutDealStateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PutDealStateResponse).Status, response.(PutDealStateResponse).Errinfo, response.(PutDealStateResponse).Data
}
