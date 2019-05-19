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

// makeGetAuthInfoHandler creates the handler logic
func makeGetAuthInfoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[a-zA-Z0-9_-]+}/certs").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(
			http.NewServer(
				endpoints.GetAuthInfoEndpoint,
				decodeGetAuthInfoRequest,
				encodeGetAuthInfoResponse,
				options...,
			),
		),
	)
}

// decodeGetAuthInfoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAuthInfoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.GetAuthInfoRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeGetAuthInfoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAuthInfoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostAuthInfoHandler creates the handler logic
func makePostAuthInfoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/users/{userId:[a-zA-Z0-9_-]+}/certs").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(
			http.NewServer(endpoints.PostAuthInfoEndpoint,
				decodePostAuthInfoRequest,
				encodePostAuthInfoResponse,
				options...,
			),
		),
	)
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

// makeGetAllUnCertifyHandler creates the handler logic
func makeGetAllUnCertifyHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/certs").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetAllUnCertifyEndpoint, decodeGetAllUnCertifyRequest, encodeGetAllUnCertifyResponse, options...)))
}

// decodeGetAllUnCertifyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAllUnCertifyRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetAllUnCertifyRequest{}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeGetAllUnCertifyResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAllUnCertifyResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUnCertifyInfoHandler creates the handler logic
func makeGetUnCertifyInfoHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/certs/{userId:[a-zA-Z0-9_-]+}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetUnCertifyInfoEndpoint, decodeGetUnCertifyInfoRequest, encodeGetUnCertifyInfoResponse, options...)))
}

// decodeGetUnCertifyInfoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUnCertifyInfoRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.GetUnCertifyInfoRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeGetUnCertifyInfoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUnCertifyInfoResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostCertifyStateHandler creates the handler logic
func makePostCertifyStateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/certs/{userId:[a-zA-Z0-9_-]+}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostCertifyStateEndpoint, decodePostCertifyStateRequest, encodePostCertifyStateResponse, options...)))
}

// decodePostCertifyStateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostCertifyStateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.PostCertifyStateRequest{
		Id: id,
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostCertifyStateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostCertifyStateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
