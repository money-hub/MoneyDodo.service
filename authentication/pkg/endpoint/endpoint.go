package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/money-hub/MoneyDodo.service/authentication/pkg/service"
)

// GetOpenidRequest collects the request parameters for the GetOpenid method.
type GetOpenidRequest struct {
	Code string `json:"code"`
}

// GetOpenidResponse collects the response parameters for the GetOpenid method.
type GetOpenidResponse struct {
	Status  bool   `json:"status"`
	Errinfo string `json:"errinfo"`
	Data    string `json:"data"`
}

// MakeGetOpenidEndpoint returns an endpoint that invokes GetOpenid on the service.
func MakeGetOpenidEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOpenidRequest)
		status, errinfo, data := s.GetOpenid(ctx, req.Code)
		return GetOpenidResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// AdminLoginRequest collects the request parameters for the AdminLogin method.
type AdminLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// AdminLoginResponse collects the response parameters for the AdminLogin method.
type AdminLoginResponse struct {
	Status  bool   `json:"status"`
	Errinfo string `json:"errinfo"`
	Data    string `json:"data"`
}

// MakeAdminLoginEndpoint returns an endpoint that invokes AdminLogin on the service.
func MakeAdminLoginEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AdminLoginRequest)
		status, errinfo, data := s.AdminLogin(ctx, req.Name, req.Password)
		return AdminLoginResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetOpenid implements Service. Primarily useful in a client.
func (e Endpoints) GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data string) {
	request := GetOpenidRequest{Code: code}
	response, err := e.GetOpenidEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetOpenidResponse).Status, response.(GetOpenidResponse).Errinfo, response.(GetOpenidResponse).Data
}

// AdminLogin implements Service. Primarily useful in a client.
func (e Endpoints) AdminLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string) {
	request := AdminLoginRequest{
		Name:     name,
		Password: password,
	}
	response, err := e.AdminLoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AdminLoginResponse).Status, response.(AdminLoginResponse).Errinfo, response.(AdminLoginResponse).Data
}
