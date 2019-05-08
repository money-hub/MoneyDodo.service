package http

import (
	"context"
	"encoding/json"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/comment/pkg/endpoint"
	http1 "net/http"
)

// makeGetCommentHandler creates the handler logic
func makeGetCommentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/get-comment").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetCommentEndpoint, decodeGetCommentRequest, encodeGetCommentResponse, options...)))
}

// decodeGetCommentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetCommentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetCommentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetCommentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostCommentHandler creates the handler logic
func makePostCommentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/post-comment").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostCommentEndpoint, decodePostCommentRequest, encodePostCommentResponse, options...)))
}

// decodePostCommentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostCommentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PostCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostCommentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostCommentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeChangeCommentHandler creates the handler logic
func makeChangeCommentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/change-comment").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ChangeCommentEndpoint, decodeChangeCommentRequest, encodeChangeCommentResponse, options...)))
}

// decodeChangeCommentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeChangeCommentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ChangeCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeChangeCommentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeChangeCommentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteCommentHandler creates the handler logic
func makeDeleteCommentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/delete-comment").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteCommentEndpoint, decodeDeleteCommentRequest, encodeDeleteCommentResponse, options...)))
}

// decodeDeleteCommentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteCommentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteCommentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteCommentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeLikeCommentHandler creates the handler logic
func makeLikeCommentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/like-comment").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.LikeCommentEndpoint, decodeLikeCommentRequest, encodeLikeCommentResponse, options...)))
}

// decodeLikeCommentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeLikeCommentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.LikeCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeLikeCommentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeLikeCommentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCancelLikeCommentHandler creates the handler logic
func makeCancelLikeCommentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/cancel-like-comment").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CancelLikeCommentEndpoint, decodeCancelLikeCommentRequest, encodeCancelLikeCommentResponse, options...)))
}

// decodeCancelLikeCommentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCancelLikeCommentRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CancelLikeCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCancelLikeCommentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCancelLikeCommentResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
