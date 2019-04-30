package cpt

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/model"
)

type Decodes struct {
	GetSpecDecode func(ctx context.Context, r *http.Request) (interface{}, error)
	GetAllDecode  func(ctx context.Context, r *http.Request) (interface{}, error)
	PostDecode    func(ctx context.Context, r *http.Request) (interface{}, error)
	PutDecode     func(ctx context.Context, r *http.Request) (interface{}, error)
	DeleteDecode  func(ctx context.Context, r *http.Request) (interface{}, error)
}

type Encodes struct {
	GetSpecEncode func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	GetAllEncode  func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	PostEncode    func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	PutEncode     func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	DeleteEncode  func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
}

func MakeServerDecodes() Decodes {
	return Decodes{
		GetSpecDecode: decodeRequest,
		GetAllDecode:  decodeRequest,
		PostDecode:    decodeRequest,
		PutDecode:     decodeRequest,
		DeleteDecode:  decodeRequest,
	}
}

func MakeServerEncodes() Encodes {
	return Encodes{
		GetSpecEncode: encodeResponse,
		GetAllEncode:  encodeResponse,
		PostEncode:    encodeResponse,
		PutEncode:     encodeResponse,
		DeleteEncode:  encodeResponse,
	}
}

// decodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	taskId, ok := vars["taskId"]
	if !ok {
		log.Println("taskId is not in the request URL.")
	}
	state, ok := vars["state"]
	if !ok {
		log.Println("state is not in the request URL.")
	}
	req := Request{
		TaskId: taskId,
		State:  state,
	}
	if r.Method == "POST" || r.Method == "PUT" {
		wrapper := model.Wrapper{}
		err := json.NewDecoder(r.Body).Decode(&wrapper)
		if err != nil {
			log.Println("The upload wrapper is not correct.")
		}
		req.Kind = wrapper.Kind
		if req.Kind == model.TaskKindQuestionnaire {
			t := model.Qtnr{
				Qtnr: &model.Questionnaire{},
			}
			err := json.Unmarshal(wrapper.Raw, &t)
			if err != nil {
				log.Println("task is not in the request Body.")
			}
			req.Task = t
		}
	}
	return req, nil
}

// encodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// // decodeGetAllRequest is a transport/http.DecodeRequestFunc that decodes a
// // JSON-encoded requestfrom the HTTP request body.
// func decodeGetAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	userId, ok := vars["userId"]
// 	if !ok {
// 		return nil, errors.New("not a valid userId")
// 	}
// 	req := GetAllRequest{
// 		UserId: userId,
// 	}
// 	return req, nil
// }

// // decodePostRequest is a transport/http.DecodeRequestFunc that decodes a
// // JSON-encoded requestfrom the HTTP request body.
// func decodePostRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	userId, ok := vars["userId"]
// 	if !ok {
// 		return nil, errors.New("not a valid userId")
// 	}
// 	req := PostRequest{
// 		UserId: userId,
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&req.Task)
// 	return req, err
// }

// // decodePostClaimRequest is a transport/http.DecodeRequestFunc that decodes a
// // JSON-encoded requestfrom the HTTP request body.
// func decodePostClaimRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	userId, ok1 := vars["userId"]
// 	taskId, ok2 := vars["taskId"]
// 	if !ok1 {
// 		return nil, errors.New("not a valid userId")
// 	}
// 	if !ok2 {
// 		return nil, errors.New("not a valid taskId")
// 	}
// 	req := PostClaimRequest{
// 		UserId: userId,
// 		TaskId: taskId,
// 	}
// 	return req, nil
// }

// // decodePutRequest is a transport/http.DecodeRequestFunc that decodes a
// // JSON-encoded requestfrom the HTTP request body.
// func decodePutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	userId, ok := vars["userId"]
// 	taskId, ok2 := vars["taskId"]
// 	if !ok {
// 		return nil, errors.New("not a valid userId")
// 	}
// 	if !ok2 {
// 		return nil, errors.New("not a valid taskId")
// 	}
// 	req := PutRequest{
// 		UserId: userId,
// 		TaskId: taskId,
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&req.Task)
// 	return req, err
// }

// // decodeDeleteRequest is a transport/http.DecodeRequestFunc that decodes a
// // JSON-encoded requestfrom the HTTP request body.
// func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	userId, ok1 := vars["userId"]
// 	taskId, ok2 := vars["taskId"]
// 	status, ok3 := vars["status"]
// 	if !ok1 {
// 		return nil, errors.New("not a valid userId")
// 	}
// 	if !ok2 {
// 		return nil, errors.New("not a valid taskId")
// 	}
// 	if !ok3 {
// 		return nil, errors.New("not a valid status")
// 	}
// 	req := DeleteRequest{
// 		UserId: userId,
// 		TaskId: taskId,
// 		Status: status,
// 	}
// 	return req, nil
// }
