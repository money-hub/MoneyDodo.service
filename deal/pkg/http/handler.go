package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/deal/pkg/endpoint"
)

// makeGetUserDealByStateHandler creates the handler logic
func makeGetUserDealByStateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[a-zA-Z0-9_-]+}/deals").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetUserDealByStateEndpoint, decodeGetUserDealByStateRequest, encodeGetUserDealByStateResponse, options...)))
}

// decodeGetUserDealByStateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserDealByStateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	vals := r.URL.Query()
	state := ""
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	states, okk := vals["state"]
	if okk {
		state = states[0]
	}
	req := endpoint.GetUserDealByStateRequest{
		Id:    id,
		State: state,
	}
	return req, nil
}

// encodeGetUserDealByStateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserDealByStateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetDealByDIDHandler creates the handler logic
func makeGetDealByDIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/deals/{dId:[0-9]+}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetDealByDIDEndpoint, decodeGetDealByDIDRequest, encodeGetDealByDIDResponse, options...)))
}

// decodeGetDealByDIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetDealByDIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["dId"]
	if !ok {
		return nil, errors.New("not a valid dId")
	}
	req := endpoint.GetDealByDIDRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetDealByDIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetDealByDIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetDealByStateHandler creates the handler logic
func makeGetDealByStateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/deals").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetDealByStateEndpoint, decodeGetDealByStateRequest, encodeGetDealByStateResponse, options...)))
}

// decodeGetDealByStateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetDealByStateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vals := r.URL.Query()
	state := ""
	states, okk := vals["state"]
	if okk {
		state = states[0]
	}
	req := endpoint.GetDealByStateRequest{
		State: state,
	}
	return req, nil
}

// encodeGetDealByStateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetDealByStateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostAcceptDealHandler creates the handler logic
func makePostAcceptDealHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/deals").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostAcceptDealEndpoint, decodePostAcceptDealRequest, encodePostAcceptDealResponse, options...)))
}

// decodePostAcceptDealRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostAcceptDealRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PostAcceptDealRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostAcceptDealResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostAcceptDealResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
