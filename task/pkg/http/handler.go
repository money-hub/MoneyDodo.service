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

// makeGetTasksByIDHandler creates the handler logic
func makeGetTasksByIDHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/users/{userId:[a-zA-Z0-9_-]+}/tasks").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetTasksByIDEndpoint, decodeGetTasksByIDRequest, encodeGetTasksByIDResponse, options...)))
}

// decodeGetTasksByIDRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTasksByIDRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	vals := r.URL.Query()
	state := ""
	id, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	states, ok := vals["state"]
	if ok {
		state = states[0]
	} else {
		return nil, errors.New("not a valid state")
	}
	req := endpoint.GetTasksByIDRequest{
		Id:    id,
		State: state,
	}
	return req, nil
}

// encodeGetTasksByIDResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTasksByIDResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
