package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/money-hub/MoneyDodo.service/authentication/pkg/service"
)

// GetOpenIdRequest collects the request parameters for the GetOpenId method.
type GetOpenIdRequest struct {
	Code string `json:"code"`
}

// GetOpenIdResponse collects the response parameters for the GetOpenId method.
type GetOpenIdResponse struct {
	E0 error  `json:"e0"`
	S1 string `json:"s1"`
	S2 string `json:"s2"`
}

// MakeGetOpenIdEndpoint returns an endpoint that invokes GetOpenId on the service.
func MakeGetOpenIdEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOpenIdRequest)
		e0, s1, s2 := s.GetOpenId(ctx, req.Code)
		return GetOpenIdResponse{
			E0: e0,
			S1: s1,
			S2: s2,
		}, nil
	}
}

// Failed implements Failer.
func (r GetOpenIdResponse) Failed() error {
	return r.E0
}

// AdminLoginRequest collects the request parameters for the AdminLogin method.
type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminLoginResponse collects the response parameters for the AdminLogin method.
type AdminLoginResponse struct {
	E0 error  `json:"e0"`
	B1 bool   `json:"b1"`
	S2 string `json:"s2"`
}

// MakeAdminLoginEndpoint returns an endpoint that invokes AdminLogin on the service.
func MakeAdminLoginEndpoint(s service.AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AdminLoginRequest)
		e0, b1, s2 := s.AdminLogin(ctx, req.Username, req.Password)
		return AdminLoginResponse{
			B1: b1,
			E0: e0,
			S2: s2,
		}, nil
	}
}

// Failed implements Failer.
func (r AdminLoginResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetOpenId implements Service. Primarily useful in a client.
func (e Endpoints) GetOpenId(ctx context.Context, code string) (e0 error, s1 string, s2 string) {
	request := GetOpenIdRequest{Code: code}
	response, err := e.GetOpenIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetOpenIdResponse).E0, response.(GetOpenIdResponse).S1, response.(GetOpenIdResponse).S2
}

// AdminLogin implements Service. Primarily useful in a client.
func (e Endpoints) AdminLogin(ctx context.Context, username string, password string) (e0 error, b1 bool, s2 string) {
	request := AdminLoginRequest{
		Password: password,
		Username: username,
	}
	response, err := e.AdminLoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AdminLoginResponse).E0, response.(AdminLoginResponse).B1, response.(AdminLoginResponse).S2
}
