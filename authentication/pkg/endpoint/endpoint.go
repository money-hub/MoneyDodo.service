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
	Status  bool             `json:"status"`
	Errinfo string           `json:"errinfo"`
	Data    *service.UserRes `json:"data"`
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

// EnterpriseLoginRequest collects the request parameters for the EnterpriseLogin method.
type EnterpriseLoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// EnterpriseLoginResponse collects the response parameters for the EnterpriseLogin method.
type EnterpriseLoginResponse struct {
	Status  bool   `json:"status"`
	Errinfo string `json:"errinfo"`
	Data    string `json:"data"`
}

// MakeEnterpriseLoginEndpoint returns an endpoint that invokes EnterpriseLogin on the service.
func MakeEnterpriseLoginEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EnterpriseLoginRequest)
		status, errinfo, data := s.EnterpriseLogin(ctx, req.Name, req.Password)
		return EnterpriseLoginResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// LogoutRequest collects the request parameters for the Logout method.
type LogoutRequest struct{}

// LogoutResponse collects the response parameters for the Logout method.
type LogoutResponse struct {
	Status  bool   `json:"status"`
	Errinfo string `json:"errinfo"`
	Data    string `json:"data"`
}

// MakeLogoutEndpoint returns an endpoint that invokes Logout on the service.
func MakeLogoutEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status, errinfo, data := s.Logout(ctx)
		return LogoutResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetOpenid implements Service. Primarily useful in a client.
func (e Endpoints) GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data *service.UserRes) {
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

// EnterpriseLogin implements Service. Primarily useful in a client.
func (e Endpoints) EnterpriseLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string) {
	request := EnterpriseLoginRequest{
		Name:     name,
		Password: password,
	}
	response, err := e.EnterpriseLoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(EnterpriseLoginResponse).Status, response.(EnterpriseLoginResponse).Errinfo, response.(EnterpriseLoginResponse).Data
}

// Logout implements Service. Primarily useful in a client.
func (e Endpoints) Logout(ctx context.Context) (status bool, errinfo string, data string) {
	request := LogoutRequest{}
	response, err := e.LogoutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LogoutResponse).Status, response.(LogoutResponse).Errinfo, response.(LogoutResponse).Data
}
