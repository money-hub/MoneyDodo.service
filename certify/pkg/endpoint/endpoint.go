package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/money-hub/MoneyDodo.service/certify/pkg/service"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// GetAllUnAuthRequest collects the request parameters for the GetAllUnAuth method.
type GetAllUnAuthRequest struct{}

// GetAllUnAuthResponse collects the response parameters for the GetAllUnAuth method.
type GetAllUnAuthResponse struct {
	Status  bool         `json:"status"`
	Errinfo string       `json:"errinfo"`
	Data    []model.User `json:"data"`
}

// MakeGetAllUnAuthEndpoint returns an endpoint that invokes GetAllUnAuth on the service.
func MakeGetAllUnAuthEndpoint(s service.CertifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status, errinfo, data := s.GetAllUnAuth(ctx)
		return GetAllUnAuthResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostAuthInfoRequest collects the request parameters for the PostAuthInfo method.
type PostAuthInfoRequest struct {
	Id  string `json:"id"`
	Img []byte `json:"img"`
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
		status, errinfo, data := s.PostAuthInfo(ctx, req.Id, req.Img)
		return PostAuthInfoResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostCertifyInfoRequest collects the request parameters for the PostCertifyInfo method.
type PostCertifyInfoRequest struct {
	Id   string `json:"id"`
	Pass bool   `json:"pass"`
}

// PostCertifyInfoResponse collects the response parameters for the PostCertifyInfo method.
type PostCertifyInfoResponse struct {
	Status  bool       `json:"status"`
	Errinfo string     `json:"errinfo"`
	Data    model.User `json:"data"`
}

// MakePostCertifyInfoEndpoint returns an endpoint that invokes PostCertifyInfo on the service.
func MakePostCertifyInfoEndpoint(s service.CertifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCertifyInfoRequest)
		status, errinfo, data := s.PostCertifyInfo(ctx, req.Id, req.Pass)
		return PostCertifyInfoResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetAllUnAuth implements Service. Primarily useful in a client.
func (e Endpoints) GetAllUnAuth(ctx context.Context) (status bool, errinfo string, data []model.User) {
	request := GetAllUnAuthRequest{}
	response, err := e.GetAllUnAuthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllUnAuthResponse).Status, response.(GetAllUnAuthResponse).Errinfo, response.(GetAllUnAuthResponse).Data
}

// PostAuthInfo implements Service. Primarily useful in a client.
func (e Endpoints) PostAuthInfo(ctx context.Context, id string, img []byte) (status bool, errinfo string, data model.User) {
	request := PostAuthInfoRequest{
		Id:  id,
		Img: img,
	}
	response, err := e.PostAuthInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostAuthInfoResponse).Status, response.(PostAuthInfoResponse).Errinfo, response.(PostAuthInfoResponse).Data
}

// PostCertifyInfo implements Service. Primarily useful in a client.
func (e Endpoints) PostCertifyInfo(ctx context.Context, id string, pass bool) (status bool, errinfo string, data model.User) {
	request := PostCertifyInfoRequest{
		Id:   id,
		Pass: pass,
	}
	response, err := e.PostCertifyInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostCertifyInfoResponse).Status, response.(PostCertifyInfoResponse).Errinfo, response.(PostCertifyInfoResponse).Data
}
