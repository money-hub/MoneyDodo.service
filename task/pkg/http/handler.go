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

// makeUserGetHisReleasedTasksHandler creates the handler logic
func makeUserGetHisReleasedTasksHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UserGetHisReleasedTasksEndpoint, decodeUserGetHisReleasedTasksRequest, encodeUserGetHisReleasedTasksResponse, options...)))
}

// decodeUserGetHisReleasedTasksRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUserGetHisReleasedTasksRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.UserGetHisReleasedTasksRequest{
		Id: id,
	}
	return req, nil
}

// encodeUserGetHisReleasedTasksResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUserGetHisReleasedTasksResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUserGetTasksByIDHandler creates the handler logic
func makeUserGetTasksByIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=released").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UserGetTasksByIDEndpoint, decodeUserGetTasksByIDRequest, encodeUserGetTasksByIDResponse, options...)))
}

// decodeUserGetTasksByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUserGetTasksByIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.UserGetTasksByIDRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeUserGetTasksByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUserGetTasksByIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUserGetHisUnreleasedTasksHandler creates the handler logic
func makeUserGetHisUnreleasedTasksHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=non-released").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UserGetHisUnreleasedTasksEndpoint, decodeUserGetHisUnreleasedTasksRequest, encodeUserGetHisUnreleasedTasksResponse, options...)))
}

// decodeUserGetHisUnreleasedTasksRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUserGetHisUnreleasedTasksRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.UserGetHisUnreleasedTasksRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeUserGetHisUnreleasedTasksResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUserGetHisUnreleasedTasksResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUserGetHisClosedTasksHandler creates the handler logic
func makeUserGetHisClosedTasksHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=closed").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UserGetHisClosedTasksEndpoint, decodeUserGetHisClosedTasksRequest, encodeUserGetHisClosedTasksResponse, options...)))
}

// decodeUserGetHisClosedTasksRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUserGetHisClosedTasksRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.UserGetHisClosedTasksRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeUserGetHisClosedTasksResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUserGetHisClosedTasksResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAdminGetAllTasksByUserIDHandler creates the handler logic
func makeAdminGetAllTasksByUserIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AdminGetAllTasksByUserIDEndpoint, decodeAdminGetAllTasksByUserIDRequest, encodeAdminGetAllTasksByUserIDResponse, options...)))
}

// decodeAdminGetAllTasksByUserIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAdminGetAllTasksByUserIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.AdminGetAllTasksByUserIDRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeAdminGetAllTasksByUserIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAdminGetAllTasksByUserIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAdminGetTasksReleasedByUserIDHandler creates the handler logic
func makeAdminGetTasksReleasedByUserIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=released").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AdminGetTasksReleasedByUserIDEndpoint, decodeAdminGetTasksReleasedByUserIDRequest, encodeAdminGetTasksReleasedByUserIDResponse, options...)))
}

// decodeAdminGetTasksReleasedByUserIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAdminGetTasksReleasedByUserIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.AdminGetTasksReleasedByUserIDRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeAdminGetTasksReleasedByUserIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAdminGetTasksReleasedByUserIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAdminGetTasksUnreleasedByUserIDHandler creates the handler logic
func makeAdminGetTasksUnreleasedByUserIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=non-released").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AdminGetTasksUnreleasedByUserIDEndpoint, decodeAdminGetTasksUnreleasedByUserIDRequest, encodeAdminGetTasksUnreleasedByUserIDResponse, options...)))
}

// decodeAdminGetTasksUnreleasedByUserIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAdminGetTasksUnreleasedByUserIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.AdminGetTasksUnreleasedByUserIDRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeAdminGetTasksUnreleasedByUserIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAdminGetTasksUnreleasedByUserIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAdminGetTasksClosedByUserIDHandler creates the handler logic
func makeAdminGetTasksClosedByUserIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[0-9]+}/tasks?state=closed").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AdminGetTasksClosedByUserIDEndpoint, decodeAdminGetTasksClosedByUserIDRequest, encodeAdminGetTasksClosedByUserIDResponse, options...)))
}

// decodeAdminGetTasksClosedByUserIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAdminGetTasksClosedByUserIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := endpoint.AdminGetTasksClosedByUserIDRequest{
		Id: id,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeAdminGetTasksClosedByUserIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAdminGetTasksClosedByUserIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
