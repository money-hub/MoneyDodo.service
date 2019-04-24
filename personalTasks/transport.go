package personalTasks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// decodeGetSpecRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodeGetSpecRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok1 := vars["userId"]
	taskId, ok2 := vars["taskId"]
	if !ok1 {
		return nil, errors.New("not a valid userId")
	}
	if !ok2 {
		return nil, errors.New("not a valid taskId")
	}
	req := GetSpecRequest{
		UserId: userId,
		TaskId: taskId,
	}
	return req, nil
}

// encodeGetSpecResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetSpecResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeGetAllRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodeGetAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := GetAllRequest{
		UserId: userId,
	}
	return req, nil
}

// encodeGetAllResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAllResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodePostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodePostRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := PostRequest{
		UserId: userId,
	}
	err := json.NewDecoder(r.Body).Decode(&req.Task)
	return req, err
}

// encodePostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodePostClaimRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodePostClaimRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok1 := vars["userId"]
	taskId, ok2 := vars["taskId"]
	if !ok1 {
		return nil, errors.New("not a valid userId")
	}
	if !ok2 {
		return nil, errors.New("not a valid taskId")
	}
	req := PostClaimRequest{
		UserId: userId,
		TaskId: taskId,
	}
	return req, nil
}

// encodePostClaimResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostClaimResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodePutRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodePutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	req := PutRequest{
		UserId: userId,
	}
	err := json.NewDecoder(r.Body).Decode(&req.Task)
	return req, err
}

// encodePutResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePutResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// decodeDeleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok1 := vars["userId"]
	taskId, ok2 := vars["taskId"]
	detail, ok3 := vars["detail"]
	if !ok1 {
		return nil, errors.New("not a valid userId")
	}
	if !ok2 {
		return nil, errors.New("not a valid taskId")
	}
	if !ok3 {
		return nil, errors.New("not a valid detail")
	}
	req := DeleteRequest{
		UserId: userId,
		TaskId: taskId,
		Detail: detail,
	}
	return req, nil
}

// encodeDeleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
