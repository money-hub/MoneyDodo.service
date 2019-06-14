package http

import (
	"context"
	"encoding/json"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/comment/pkg/endpoint"
)

// makeGetCommentHandler creates the handler logic
func makeGetCommentHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation GET /api/tasks/{taskId}/comment Comment swaggGetCommentReq
	// ---
	// summary: Get a comment from a task
	// description: Get a comment from a task
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggCommentsResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("GET").Path("/api/tasks/{taskId:[0-9]+}/comments").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetCommentEndpoint, decodeGetCommentRequest, encodeGetCommentResponse, options...)))
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
	// swagger:operation POST /api/tasks/{taskId}/comments Comment swaggCommentReq
	// ---
	// summary: Post a comment
	// description: You need to specify comment information to post a comment.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: comment
	//   in: path
	//   description: comment's content
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggCommentResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("POST").Path("/api/tasks/{taskId:[0-9]+}/comments").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostCommentEndpoint, decodePostCommentRequest, encodePostCommentResponse, options...)))
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
	// swagger:operation PUT /api/tasks/{taskId}/comments/{cid} Comment swaggCommentReq
	// ---
	// summary: Change the comment
	// description: You need to specify comment information to change a comment and only the reviewer himself can edit the comment.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: cId
	//   in: path
	//   description: id of comment
	//   type: string
	//   required: true
	// - name: comment
	//   in: path
	//   description: comment's content
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggCommentResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("PUT").Path("/api/tasks/{taskId:[0-9]+}/comments/{cid:[0-9]+}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ChangeCommentEndpoint, decodeChangeCommentRequest, encodeChangeCommentResponse, options...)))
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
	// swagger:operation DELETE /api/tasks/{taskId}/comments/{cid} authentication
	// ---
	// summary: Delete the comment
	// description: You need to specify comment information to delete the comment and only the person making the comment, the task publisher, and the administrator can delete.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: cId
	//   in: path
	//   description: id of comment
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("DELETE").Path("/api/tasks/{taskId:[0-9]+}/comments/{cid:[0-9]+}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"DELETE"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteCommentEndpoint, decodeDeleteCommentRequest, encodeDeleteCommentResponse, options...)))
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
	// swagger:operation PUT /api/tasks/{taskId}/comments/{cid}/star Comment swaggStarCommentReq
	// ---
	// summary: Star the comment
	// description: You need to specify comment information to change the star status
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: cId
	//   in: path
	//   description: id of comment
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggCommentResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("PUT").Path("/api/tasks/{taskId:[0-9]+}/comments/{cid:[0-9]+}/star").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.LikeCommentEndpoint, decodeLikeCommentRequest, encodeLikeCommentResponse, options...)))
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
	// swagger:operation DELETE /api/tasks/{taskId}/comments/{cid}/star Comment swaggStarCommentReq
	// ---
	// summary: Unstar the comment
	// description: You need to specify comment information to change the star status
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: cId
	//   in: path
	//   description: id of comment
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggCommentResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("DELETE").Path("/api/tasks/{taskId:[0-9]+}/comments/{cid:[0-9]+}/star").Handler(handlers.CORS(handlers.AllowedMethods([]string{"DELETE"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CancelLikeCommentEndpoint, decodeCancelLikeCommentRequest, encodeCancelLikeCommentResponse, options...)))
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
