package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/task/pkg/endpoint"
)

// makeGetHisReleasedTasksHandler creates the handler logic
func makeGetHisReleasedTasksHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetHisReleasedTasksEndpoint, decodeGetHisReleasedTasksRequest, encodeGetHisReleasedTasksResponse, options...)))
}

// decodeGetHisReleasedTasksRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetHisReleasedTasksRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.GetHisReleasedTasksRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetHisReleasedTasksResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetHisReleasedTasksResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTasksByIDHandler creates the handler logic
func makeGetTasksByIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=released").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetTasksByIDEndpoint, decodeGetTasksByIDRequest, encodeGetTasksByIDResponse, options...)))
}

// decodeGetTasksByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTasksByIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.GetTasksByIDRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeGetTasksByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTasksByIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetHisUnreleasedTasksHandler creates the handler logic
func makeGetHisUnreleasedTasksHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=non-released").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetHisUnreleasedTasksEndpoint, decodeGetHisUnreleasedTasksRequest, encodeGetHisUnreleasedTasksResponse, options...)))
}

// decodeGetHisUnreleasedTasksRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetHisUnreleasedTasksRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.GetHisUnreleasedTasksRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeGetHisUnreleasedTasksResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetHisUnreleasedTasksResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetHisClosedTasksHandler creates the handler logic
func makeGetHisClosedTasksHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=closed").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetHisClosedTasksEndpoint, decodeGetHisClosedTasksRequest, encodeGetHisClosedTasksResponse, options...)))
}

// decodeGetHisClosedTasksRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetHisClosedTasksRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.GetHisClosedTasksRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeGetHisClosedTasksResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetHisClosedTasksResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
