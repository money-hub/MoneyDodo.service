package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/money-hub/MoneyDodo.service/certify/pkg/service"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// GetAuthInfoRequest collects the request parameters for the GetAuthInfo method.
type GetAuthInfoRequest struct {
	Id string `json:"id"`
}

// GetAuthInfoResponse collects the response parameters for the GetAuthInfo method.
type GetAuthInfoResponse struct {
	Status  bool       `json:"status"`
	Errinfo string     `json:"errinfo"`
	Data    model.User `json:"data"`
}

// MakeGetAuthInfoEndpoint returns an endpoint that invokes GetAuthInfo on the service.
func MakeGetAuthInfoEndpoint(s service.CertifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAuthInfoRequest)
		status, errinfo, data := s.GetAuthInfo(ctx, req.Id)
		return GetAuthInfoResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostAuthInfoRequest collects the request parameters for the PostAuthInfo method.
type PostAuthInfoRequest struct {
	Id           string `json:"id"`
	CertifiedPic []byte `json:"certifiedPic"`
}

// PostAuthInfoResponse collects the response parameters for the PostAuthInfo method.
type PostAuthInfoResponse struct {
	Status  bool       `json:"status"`
	Errinfo string     `json:"errinfo"`
	Data    model.User `json:"data"`
}

// MakePostAuthInfoEndpoint returns an endpoint that invokes PostAuthInfo on the service.
func MakePostAuthInfoEndpoint(s service.CertifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostAuthInfoRequest)
		status, errinfo, data := s.PostAuthInfo(ctx, req.Id, req.CertifiedPic)
		return PostAuthInfoResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetAllUnCertifyRequest collects the request parameters for the GetAllUnCertify method.
type GetAllUnCertifyRequest struct{}

// GetAllUnCertifyResponse collects the response parameters for the GetAllUnCertify method.
type GetAllUnCertifyResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.User `json:"data"`
}

// MakeGetAllUnCertifyEndpoint returns an endpoint that invokes GetAllUnCertify on the service.
func MakeGetAllUnCertifyEndpoint(s service.CertifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status, errinfo, data := s.GetAllUnCertify(ctx)
		return GetAllUnCertifyResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetUnCertifyInfoRequest collects the request parameters for the GetUnCertifyInfo method.
type GetUnCertifyInfoRequest struct {
	Id string `json:"id"`
}

// GetUnCertifyInfoResponse collects the response parameters for the GetUnCertifyInfo method.
type GetUnCertifyInfoResponse struct {
	Status  bool       `json:"status"`
	Errinfo string     `json:"errinfo"`
	Data    model.User `json:"data"`
}

// MakeGetUnCertifyInfoEndpoint returns an endpoint that invokes GetUnCertifyInfo on the service.
func MakeGetUnCertifyInfoEndpoint(s service.CertifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUnCertifyInfoRequest)
		status, errinfo, data := s.GetUnCertifyInfo(ctx, req.Id)
		return GetUnCertifyInfoResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostCertifyStateRequest collects the request parameters for the PostCertifyState method.
type PostCertifyStateRequest struct {
	Id   string `json:"id"`
	Pass bool   `json:"pass"`
}

// PostCertifyStateResponse collects the response parameters for the PostCertifyState method.
type PostCertifyStateResponse struct {
	Status  bool       `json:"status"`
	Errinfo string     `json:"errinfo"`
	Data    model.User `json:"data"`
}

// MakePostCertifyStateEndpoint returns an endpoint that invokes PostCertifyState on the service.
func MakePostCertifyStateEndpoint(s service.CertifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCertifyStateRequest)
		status, errinfo, data := s.PostCertifyState(ctx, req.Id, req.Pass)
		return PostCertifyStateResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetAuthInfo implements Service. Primarily useful in a client.
func (e Endpoints) GetAuthInfo(ctx context.Context, id string) (status bool, errinfo string, data model.User) {
	request := GetAuthInfoRequest{Id: id}
	response, err := e.GetAuthInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAuthInfoResponse).Status, response.(GetAuthInfoResponse).Errinfo, response.(GetAuthInfoResponse).Data
}

// PostAuthInfo implements Service. Primarily useful in a client.
func (e Endpoints) PostAuthInfo(ctx context.Context, id string, certifiedPic []byte) (status bool, errinfo string, data model.User) {
	request := PostAuthInfoRequest{
		Id:           id,
		CertifiedPic: certifiedPic,
	}
	response, err := e.PostAuthInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostAuthInfoResponse).Status, response.(PostAuthInfoResponse).Errinfo, response.(PostAuthInfoResponse).Data
}

// GetAllUnCertify implements Service. Primarily useful in a client.
func (e Endpoints) GetAllUnCertify(ctx context.Context) (status bool, errinfo string, data []model.User) {
	request := GetAllUnCertifyRequest{}
	response, err := e.GetAllUnCertifyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllUnCertifyResponse).Status, response.(GetAllUnCertifyResponse).Errinfo, response.(GetAllUnCertifyResponse).Data
}

// GetUnCertifyInfo implements Service. Primarily useful in a client.
func (e Endpoints) GetUnCertifyInfo(ctx context.Context, id string) (status bool, errinfo string, data model.User) {
	request := GetUnCertifyInfoRequest{Id: id}
	response, err := e.GetUnCertifyInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUnCertifyInfoResponse).Status, response.(GetUnCertifyInfoResponse).Errinfo, response.(GetUnCertifyInfoResponse).Data
}

// PostCertifyState implements Service. Primarily useful in a client.
func (e Endpoints) PostCertifyState(ctx context.Context, id string, pass bool) (status bool, errinfo string, data model.User) {
	request := PostCertifyStateRequest{
		Id:   id,
		Pass: pass,
	}
	response, err := e.PostCertifyStateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostCertifyStateResponse).Status, response.(PostCertifyStateResponse).Errinfo, response.(PostCertifyStateResponse).Data
}
