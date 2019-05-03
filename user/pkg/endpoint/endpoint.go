package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	model "github.com/money-hub/MoneyDodo.service/model"
	service "github.com/money-hub/MoneyDodo.service/user/pkg/service"
)

// GetSpecRequest collects the request parameters for the GetSpec method.
type GetSpecRequest struct {
	Id string `json:"id"`
}

// GetSpecResponse collects the response parameters for the GetSpec method.
type GetSpecResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakeGetSpecEndpoint returns an endpoint that invokes GetSpec on the service.
func MakeGetSpecEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSpecRequest)
		status, errinfo, data := s.GetSpec(ctx, req.Id)
		return GetSpecResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetAllRequest collects the request parameters for the GetAll method.
type GetAllRequest struct {
	Page    int    `json:"page"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Orderby string `json:"orderby"`
}

// GetAllResponse collects the response parameters for the GetAll method.
type GetAllResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.User `json:"data"`
}

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRequest)
		status, errinfo, data := s.GetAll(ctx, req.Page, req.Offset, req.Limit, req.Orderby)
		return GetAllResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetUDFRequest collects the request parameters for the GetUDF method.
type GetUDFRequest struct {
	Name    string `json:"name"`
	Page    int    `json:"page"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Orderby string `json:"orderby"`
}

// GetUDFResponse collects the response parameters for the GetUDF method.
type GetUDFResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.User `json:"data"`
}

// MakeGetUDFEndpoint returns an endpoint that invokes GetUDF on the service.
func MakeGetUDFEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUDFRequest)
		status, errinfo, data := s.GetUDF(ctx, req.Name, req.Page, req.Offset, req.Limit, req.Orderby)
		return GetUDFResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostRequest collects the request parameters for the Post method.
type PostRequest struct {
	User model.User `json:"user"`
}

// PostResponse collects the response parameters for the Post method.
type PostResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		status, errinfo, data := s.Post(ctx, req.User)
		return PostResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PatchRequest collects the request parameters for the Patch method.
type PatchRequest struct {
	Id   string     `json:"id"`
	User model.User `json:"user"`
}

// PatchResponse collects the response parameters for the Patch method.
type PatchResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakePatchEndpoint returns an endpoint that invokes Patch on the service.
func MakePatchEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PatchRequest)
		status, errinfo, data := s.Patch(ctx, req.Id, req.User)
		return PatchResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PutRequest collects the request parameters for the Put method.
type PutRequest struct {
	Id   string     `json:"id"`
	User model.User `json:"user"`
}

// PutResponse collects the response parameters for the Put method.
type PutResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakePutEndpoint returns an endpoint that invokes Put on the service.
func MakePutEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutRequest)
		status, errinfo, data := s.Put(ctx, req.Id, req.User)
		return PutResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		status, errinfo, data := s.Delete(ctx, req.Id)
		return DeleteResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetSpec implements Service. Primarily useful in a client.
func (e Endpoints) GetSpec(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	request := GetSpecRequest{Id: id}
	response, err := e.GetSpecEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetSpecResponse).Status, response.(GetSpecResponse).Errinfo, response.(GetSpecResponse).Data
}

// GetAll implements Service. Primarily useful in a client.
func (e Endpoints) GetAll(ctx context.Context) (status bool, errinfo string, data []model.User) {
	request := GetAllRequest{}
	response, err := e.GetAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllResponse).Status, response.(GetAllResponse).Errinfo, response.(GetAllResponse).Data
}

// GetUDF implements Service. Primarily useful in a client.
func (e Endpoints) GetUDF(ctx context.Context, name string) (status bool, errinfo string, data []model.User) {
	request := GetUDFRequest{Name: name}
	response, err := e.GetUDFEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUDFResponse).Status, response.(GetUDFResponse).Errinfo, response.(GetUDFResponse).Data
}

// Post implements Service. Primarily useful in a client.
func (e Endpoints) Post(ctx context.Context, user model.User) (status bool, errinfo string, data *model.User) {
	request := PostRequest{User: user}
	response, err := e.PostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostResponse).Status, response.(PostResponse).Errinfo, response.(PostResponse).Data
}

// Patch implements Service. Primarily useful in a client.
func (e Endpoints) Patch(ctx context.Context, id string, user model.User) (status bool, errinfo string, data *model.User) {
	request := PatchRequest{
		Id:   id,
		User: user,
	}
	response, err := e.PatchEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PatchResponse).Status, response.(PatchResponse).Errinfo, response.(PatchResponse).Data
}

// Put implements Service. Primarily useful in a client.
func (e Endpoints) Put(ctx context.Context, id string, user model.User) (status bool, errinfo string, data *model.User) {
	request := PutRequest{
		Id:   id,
		User: user,
	}
	response, err := e.PutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PutResponse).Status, response.(PutResponse).Errinfo, response.(PutResponse).Data
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Status, response.(DeleteResponse).Errinfo, response.(DeleteResponse).Data
}
