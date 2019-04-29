package personalTasks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type decodeFunc func(ctx context.Context, r *http.Request) (interface{}, error)
type encodeFunc func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)

type Decodes struct {
	GetSpecDecode   func(ctx context.Context, r *http.Request) (interface{}, error)
	GetAllDecode    func(ctx context.Context, r *http.Request) (interface{}, error)
	PostClaimDecode func(ctx context.Context, r *http.Request) (interface{}, error)
	PostDecode      func(ctx context.Context, r *http.Request) (interface{}, error)
	PutDecode       func(ctx context.Context, r *http.Request) (interface{}, error)
	DeleteDecode    func(ctx context.Context, r *http.Request) (interface{}, error)
}

type Encodes struct {
	GetSpecEncode   func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	GetAllEncode    func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	PostClaimEncode func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	PostEncode      func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	PutEncode       func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	DeleteEncode    func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
}

func MakeServerDecodes() Decodes {
	return Decodes{
		GetSpecDecode:   decodeGetSpecRequest,
		GetAllDecode:    decodeGetAllRequest,
		PostClaimDecode: decodePostClaimRequest,
		PostDecode:      decodePostRequest,
		PutDecode:       decodePutRequest,
		DeleteDecode:    decodeDeleteRequest,
	}
}

func MakeServerEncodes() Encodes {
	return Encodes{
		GetSpecEncode:   encodeGetSpecResponse,
		GetAllEncode:    encodeGetAllResponse,
		PostClaimEncode: encodePostClaimResponse,
		PostEncode:      encodePostResponse,
		PutEncode:       encodePutResponse,
		DeleteEncode:    encodeDeleteResponse,
	}
}

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
	taskId, ok2 := vars["taskId"]
	if !ok {
		return nil, errors.New("not a valid userId")
	}
	if !ok2 {
		return nil, errors.New("not a valid taskId")
	}
	req := PutRequest{
		UserId: userId,
		TaskId: taskId,
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
	status, ok3 := vars["status"]
	if !ok1 {
		return nil, errors.New("not a valid userId")
	}
	if !ok2 {
		return nil, errors.New("not a valid taskId")
	}
	if !ok3 {
		return nil, errors.New("not a valid status")
	}
	req := DeleteRequest{
		UserId: userId,
		TaskId: taskId,
		Status: status,
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
