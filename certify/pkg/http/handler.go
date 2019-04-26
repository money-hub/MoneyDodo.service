package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/certify/pkg/endpoint"
)

// makeGetAllUnAuthHandler creates the handler logic
func makeGetAllUnAuthHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation GET /api/users/getAllUnAuth users User
	// ---
	// summary: Get all UnAuth user(certificationStatus = 1) profile.
	// description: Only adminstrator can get those information.
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggUserResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("GET").Path("/api/users/getAllUnAuth").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(
			endpoints.GetAllUnAuthEndpoint,
			decodeGetAllUnAuthRequest,
			encodeGetAllUnAuthResponse,
			options...),
		))
}

// decodeGetAllUnAuthRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAllUnAuthRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetAllUnAuthRequest{}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeGetAllUnAuthResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAllUnAuthResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostAuthInfoHandler creates the handler logic
func makePostAuthInfoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/users/{userId:[0-9]+}/authInfo").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostAuthInfoEndpoint, decodePostAuthInfoRequest, encodePostAuthInfoResponse, options...)))
}

// decodePostAuthInfoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostAuthInfoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.PostAuthInfoRequest{
		Id: id,
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostAuthInfoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostAuthInfoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostCertifyInfoHandler creates the handler logic
func makePostCertifyInfoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/users/{userId:[0-9]+}/certifyInfo").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostCertifyInfoEndpoint, decodePostCertifyInfoRequest, encodePostCertifyInfoResponse, options...)))
}

// decodePostCertifyInfoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostCertifyInfoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.PostCertifyInfoRequest{
		Id: id,
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostCertifyInfoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostCertifyInfoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
