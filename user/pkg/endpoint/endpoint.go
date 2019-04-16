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
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	I2 interface{} `json:"i2"`
}

// MakeGetSpecEndpoint returns an endpoint that invokes GetSpec on the service.
func MakeGetSpecEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSpecRequest)
		b0, e1, i2 := s.GetSpec(ctx, req.Id)
		return GetSpecResponse{
			B0: b0,
			E1: e1,
			I2: i2,
		}, nil
	}
}

// Failed implements Failer.
func (r GetSpecResponse) Failed() error {
	return r.E1
}

// GetAllRequest collects the request parameters for the GetAll method.
type GetAllRequest struct{}

// GetAllResponse collects the response parameters for the GetAll method.
type GetAllResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	I2 interface{} `json:"i2"`
}

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		b0, e1, i2 := s.GetAll(ctx)
		return GetAllResponse{
			B0: b0,
			E1: e1,
			I2: i2,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAllResponse) Failed() error {
	return r.E1
}

// GetUDFRequest collects the request parameters for the GetUDF method.
type GetUDFRequest struct {
	Name string `json:"name"`
}

// GetUDFResponse collects the response parameters for the GetUDF method.
type GetUDFResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	I2 interface{} `json:"i2"`
}

// MakeGetUDFEndpoint returns an endpoint that invokes GetUDF on the service.
func MakeGetUDFEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUDFRequest)
		b0, e1, i2 := s.GetUDF(ctx, req.Name)
		return GetUDFResponse{
			B0: b0,
			E1: e1,
			I2: i2,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUDFResponse) Failed() error {
	return r.E1
}

// PostRequest collects the request parameters for the Post method.
type PostRequest struct {
	User model.User `json:"user"`
}

// PostResponse collects the response parameters for the Post method.
type PostResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	I2 interface{} `json:"i2"`
}

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		b0, e1, i2 := s.Post(ctx, req.User)
		return PostResponse{
			B0: b0,
			E1: e1,
			I2: i2,
		}, nil
	}
}

// Failed implements Failer.
func (r PostResponse) Failed() error {
	return r.E1
}

// PatchRequest collects the request parameters for the Patch method.
type PatchRequest struct {
	User model.User `json:"user"`
}

// PatchResponse collects the response parameters for the Patch method.
type PatchResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	I2 interface{} `json:"i2"`
}

// MakePatchEndpoint returns an endpoint that invokes Patch on the service.
func MakePatchEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PatchRequest)
		b0, e1, i2 := s.Patch(ctx, req.User)
		return PatchResponse{
			B0: b0,
			E1: e1,
			I2: i2,
		}, nil
	}
}

// Failed implements Failer.
func (r PatchResponse) Failed() error {
	return r.E1
}

// PutRequest collects the request parameters for the Put method.
type PutRequest struct {
	User model.User `json:"user"`
}

// PutResponse collects the response parameters for the Put method.
type PutResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	I2 interface{} `json:"i2"`
}

// MakePutEndpoint returns an endpoint that invokes Put on the service.
func MakePutEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutRequest)
		b0, e1, i2 := s.Put(ctx, req.User)
		return PutResponse{
			B0: b0,
			E1: e1,
			I2: i2,
		}, nil
	}
}

// Failed implements Failer.
func (r PutResponse) Failed() error {
	return r.E1
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	I2 interface{} `json:"i2"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		b0, e1, i2 := s.Delete(ctx, req.Id)
		return DeleteResponse{
			B0: b0,
			E1: e1,
			I2: i2,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetSpec implements Service. Primarily useful in a client.
func (e Endpoints) GetSpec(ctx context.Context, id string) (b0 bool, e1 error, i2 interface{}) {
	request := GetSpecRequest{Id: id}
	response, err := e.GetSpecEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetSpecResponse).B0, response.(GetSpecResponse).E1, response.(GetSpecResponse).I2
}

// GetAll implements Service. Primarily useful in a client.
func (e Endpoints) GetAll(ctx context.Context) (b0 bool, e1 error, i2 interface{}) {
	request := GetAllRequest{}
	response, err := e.GetAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllResponse).B0, response.(GetAllResponse).E1, response.(GetAllResponse).I2
}

// GetUDF implements Service. Primarily useful in a client.
func (e Endpoints) GetUDF(ctx context.Context, name string) (b0 bool, e1 error, i2 interface{}) {
	request := GetUDFRequest{Name: name}
	response, err := e.GetUDFEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUDFResponse).B0, response.(GetUDFResponse).E1, response.(GetUDFResponse).I2
}

// Post implements Service. Primarily useful in a client.
func (e Endpoints) Post(ctx context.Context, user model.User) (b0 bool, e1 error, i2 interface{}) {
	request := PostRequest{User: user}
	response, err := e.PostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostResponse).B0, response.(PostResponse).E1, response.(PostResponse).I2
}

// Patch implements Service. Primarily useful in a client.
func (e Endpoints) Patch(ctx context.Context, user model.User) (b0 bool, e1 error, i2 interface{}) {
	request := PatchRequest{User: user}
	response, err := e.PatchEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PatchResponse).B0, response.(PatchResponse).E1, response.(PatchResponse).I2
}

// Put implements Service. Primarily useful in a client.
func (e Endpoints) Put(ctx context.Context, user model.User) (b0 bool, e1 error, i2 interface{}) {
	request := PutRequest{User: user}
	response, err := e.PutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PutResponse).B0, response.(PutResponse).E1, response.(PutResponse).I2
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (b0 bool, e1 error, i2 interface{}) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).B0, response.(DeleteResponse).E1, response.(DeleteResponse).I2
}
