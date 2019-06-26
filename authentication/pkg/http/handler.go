package http

import (
	"context"
	"encoding/json"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/authentication/pkg/endpoint"
)

// makeGetOpenidHandler creates the handler logic
func makeGetOpenidHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation POST /api/auth/user authentication swaggAuthOfUserReq
	// ---
	// summary: Get the user's token with user's code.
	// description: You need to specify the code of the WeChat applet to log in to get the corresponding token.
	// parameters:
	// - name: code
	//   in: path
	//   description: code of user
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggAuthOfUserResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("POST").Path("/api/auth/user").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetOpenidEndpoint, decodeGetOpenidRequest, encodeGetOpenidResponse, options...)))
}

// decodeGetOpenidRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetOpenidRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetOpenidRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetOpenidResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetOpenidResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAdminLoginHandler creates the handler logic
func makeAdminLoginHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation POST /api/auth/admin authentication swaggAuthOfAdminReq
	// ---
	// summary: Administrator login verification
	// description: Administrator login verification
	// parameters:
	// - name: name
	//   in: path
	//   description: admin name
	//   type: string
	//   required: true
	// - name: password
	//   in: path
	//   description: admin password
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/swaggAuthOfAdminResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("POST").Path("/api/auth/admin").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AdminLoginEndpoint, decodeAdminLoginRequest, encodeAdminLoginResponse, options...)))
}

// decodeAdminLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAdminLoginRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.AdminLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAdminLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAdminLoginResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeEnterpriseLoginHandler creates the handler logic
func makeEnterpriseLoginHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/auth/firm").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.EnterpriseLoginEndpoint, decodeEnterpriseLoginRequest, encodeEnterpriseLoginResponse, options...)))
}

// decodeEnterpriseLoginRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeEnterpriseLoginRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.EnterpriseLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeEnterpriseLoginResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeEnterpriseLoginResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeLogoutHandler creates the handler logic
func makeLogoutHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation Get /api/auth/logout authentication swaggAuthLogoutReq
	// ---
	// summary: Logout
	// description: Logout
	// responses:
	//   "200":
	//     "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("Get").Path("/api/auth/logout").Handler(handlers.CORS(
		handlers.AllowedMethods([]string{"Get"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(http.NewServer(endpoints.LogoutEndpoint, decodeLogoutRequest, encodeLogoutResponse, options...)))
}

// decodeLogoutRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLogoutRequest(ctx context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.LogoutRequest{}
	// err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeLogoutResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLogoutResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
