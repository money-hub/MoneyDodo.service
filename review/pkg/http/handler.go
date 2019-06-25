package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/review/pkg/endpoint"
)

// makePostReviewHandler creates the handler logic
func makePostReviewHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/post/reviews").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostReviewEndpoint, decodePostReviewRequest, encodePostReviewResponse, options...)))
}

// decodePostReviewRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostReviewRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PostReviewRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostReviewResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostReviewResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetReviewHandler creates the handler logic
func makeGetReviewHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/reviews/{rid:[0-9]+}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetReviewEndpoint, decodeGetReviewRequest, encodeGetReviewResponse, options...)))
}

// decodeGetReviewRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetReviewRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["rid"]
	if !ok {
		return nil, errors.New("not a valid rid")
	}
	req := endpoint.GetReviewRequest{
		Rid: id,
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetReviewResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetReviewResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetReviewsHandler creates the handler logic
func makeGetReviewsHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/reviews").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetReviewsEndpoint, decodeGetReviewsRequest, encodeGetReviewsResponse, options...)))
}

// decodeGetReviewsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetReviewsRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetReviewsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetReviewsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetReviewsResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePutReviewHandler creates the handler logic
func makePutReviewHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/api/reviews/{rid}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PutReviewEndpoint, decodePutReviewRequest, encodePutReviewResponse, options...)))
}

// decodePutReviewRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePutReviewRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["rid"]
	if !ok {
		return nil, errors.New("not a valid rid")
	}
	req := endpoint.PutReviewRequest{
		Rid: id,
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePutReviewResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePutReviewResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
