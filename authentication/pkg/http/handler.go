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
	// swagger:operation POST /api/auth/user authentication swaggGetOpenidReq
	// ---
	// summary: Get the user's openid with user's code.
	// description: You need to specify the userid to get the detail profile about the user.
	// parameters:
	// - name: code
	//   in: path
	//   description: code of user
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggUserResp"
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
	// swagger:operation GET /api/users users swaggGetAllReq
	// ---
	// summary: Get all users' profiles
	// description: Get all users' profiles
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggUsersResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
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
