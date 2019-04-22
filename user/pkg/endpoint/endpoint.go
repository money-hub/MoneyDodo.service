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
	M2 *model.User `json:"m2"`
}

// MakeGetSpecEndpoint returns an endpoint that invokes GetSpec on the service.
func MakeGetSpecEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSpecRequest)
		b0, e1, m2 := s.GetSpec(ctx, req.Id)
		return GetSpecResponse{
			B0: b0,
			E1: e1,
			M2: m2,
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
	B0 bool         `json:"b0"`
	E1 error        `json:"e1"`
	M2 []model.User `json:"m2"`
}

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		b0, e1, m2 := s.GetAll(ctx)
		return GetAllResponse{
			B0: b0,
			E1: e1,
			M2: m2,
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
	B0 bool         `json:"b0"`
	E1 error        `json:"e1"`
	M2 []model.User `json:"m2"`
}

// MakeGetUDFEndpoint returns an endpoint that invokes GetUDF on the service.
func MakeGetUDFEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUDFRequest)
		b0, e1, m2 := s.GetUDF(ctx, req.Name)
		return GetUDFResponse{
			B0: b0,
			E1: e1,
			M2: m2,
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
	M2 *model.User `json:"m2"`
}

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		b0, e1, m2 := s.Post(ctx, req.User)
		return PostResponse{
			B0: b0,
			E1: e1,
			M2: m2,
		}, nil
	}
}

// Failed implements Failer.
func (r PostResponse) Failed() error {
	return r.E1
}

// PatchRequest collects the request parameters for the Patch method.
type PatchRequest struct {
	Id   string     `json:"id"`
	User model.User `json:"user"`
}

// PatchResponse collects the response parameters for the Patch method.
type PatchResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	M2 *model.User `json:"m2"`
}

// MakePatchEndpoint returns an endpoint that invokes Patch on the service.
func MakePatchEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PatchRequest)
		b0, e1, m2 := s.Patch(ctx, req.Id, req.User)
		return PatchResponse{
			B0: b0,
			E1: e1,
			M2: m2,
		}, nil
	}
}

// Failed implements Failer.
func (r PatchResponse) Failed() error {
	return r.E1
}

// PutRequest collects the request parameters for the Put method.
type PutRequest struct {
	Id   string     `json:"id"`
	User model.User `json:"user"`
}

// PutResponse collects the response parameters for the Put method.
type PutResponse struct {
	B0 bool        `json:"b0"`
	E1 error       `json:"e1"`
	M2 *model.User `json:"m2"`
}

// MakePutEndpoint returns an endpoint that invokes Put on the service.
func MakePutEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutRequest)
		b0, e1, m2 := s.Put(ctx, req.Id, req.User)
		return PutResponse{
			B0: b0,
			E1: e1,
			M2: m2,
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
	M2 *model.User `json:"m2"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		b0, e1, m2 := s.Delete(ctx, req.Id)
		return DeleteResponse{
			B0: b0,
			E1: e1,
			M2: m2,
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
func (e Endpoints) GetSpec(ctx context.Context, id string) (b0 bool, e1 error, m2 *model.User) {
	request := GetSpecRequest{Id: id}
	response, err := e.GetSpecEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetSpecResponse).B0, response.(GetSpecResponse).E1, response.(GetSpecResponse).M2
}

// GetAll implements Service. Primarily useful in a client.
func (e Endpoints) GetAll(ctx context.Context) (b0 bool, e1 error, m2 []model.User) {
	request := GetAllRequest{}
	response, err := e.GetAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllResponse).B0, response.(GetAllResponse).E1, response.(GetAllResponse).M2
}

// GetUDF implements Service. Primarily useful in a client.
func (e Endpoints) GetUDF(ctx context.Context, name string) (b0 bool, e1 error, m2 []model.User) {
	request := GetUDFRequest{Name: name}
	response, err := e.GetUDFEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUDFResponse).B0, response.(GetUDFResponse).E1, response.(GetUDFResponse).M2
}

// Post implements Service. Primarily useful in a client.
func (e Endpoints) Post(ctx context.Context, user model.User) (b0 bool, e1 error, m2 *model.User) {
	request := PostRequest{User: user}
	response, err := e.PostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostResponse).B0, response.(PostResponse).E1, response.(PostResponse).M2
}

// Patch implements Service. Primarily useful in a client.
func (e Endpoints) Patch(ctx context.Context, id string, user model.User) (b0 bool, e1 error, m2 *model.User) {
	request := PatchRequest{
		Id:   id,
		User: user,
	}
	response, err := e.PatchEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PatchResponse).B0, response.(PatchResponse).E1, response.(PatchResponse).M2
}

// Put implements Service. Primarily useful in a client.
func (e Endpoints) Put(ctx context.Context, id string, user model.User) (b0 bool, e1 error, m2 *model.User) {
	request := PutRequest{
		Id:   id,
		User: user,
	}
	response, err := e.PutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PutResponse).B0, response.(PutResponse).E1, response.(PutResponse).M2
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (b0 bool, e1 error, m2 *model.User) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).B0, response.(DeleteResponse).E1, response.(DeleteResponse).M2
}
