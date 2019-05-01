package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	http1 "net/http"
	"strconv"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/user/pkg/endpoint"
)

func getQueries(r *http1.Request) (error, int, int, int, string) {
	vars := mux.Vars(r)
	var page, offset, limit int
	var err error
	pageStr, ok := vars["page"]
	if !ok {
		log.Println("page is not in the request URL.")
	}
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return err, 0, 0, 0, ""
	}
	offsetStr, ok := vars["offset"]
	if !ok {
		log.Println("offset is not in the request URL.")
	}
	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		return err, 0, 0, 0, ""
	}
	limitStr, ok := vars["limit"]
	if !ok {
		log.Println("limit is not in the request URL.")
	}
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return err, 0, 0, 0, ""
	}
	orderby, ok := vars["orderby"]
	if !ok {
		log.Println("orderby is not in the request URL.")
	}
	return nil, page, offset, limit, orderby
}

// makeGetSpecHandler creates the handler logic
func makeGetSpecHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation GET /api/users/{userid} users swaggGetSpecReq
	// ---
	// summary: Get the user's profile with userid.
	// description: You need to specify the userid to get the detail profile about the user.
	// parameters:
	// - name: userid
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggUserResp"
	//   "400":
	//     "$ref": "#/responses/swaggBadReq"
	m.Methods("GET").Path("/api/users/{userid:[0-9]+}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetSpecEndpoint, decodeGetSpecRequest, encodeGetSpecResponse, options...)),
	)
}

// decodeGetSpecRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetSpecRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userid"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.GetSpecRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetSpecResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetSpecResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAllHandler creates the handler logic
func makeGetAllHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation GET /api/users users swaggGetAllReq
	// ---
	// summary: Get all users' profiles
	// description: Get all users' profiles
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggUsersResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	m.Methods("GET").Path("/api/users").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetAllEndpoint, decodeGetAllRequest, encodeGetAllResponse, options...)),
	).Queries("page", "{page}", "offset", "{offset}", "limit", "{limit}", "orderby", "{orderby}")
}

// decodeGetAllRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAllRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	err, page, offset, limit, orderby := getQueries(r)
	req := endpoint.GetAllRequest{
		Page:    page,
		Offset:  offset,
		Limit:   limit,
		Orderby: orderby,
	}
	return req, err
}

// encodeGetAllResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAllResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUDFHandler creates the handler logic
func makeGetUDFHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation GET /api/users/{username} users swaggGetUDFReq
	// ---
	// summary: Get all users' profiles with the request username
	// description: Get all users' profiles with the request username
	// parameters:
	// - name: username
	//   in: path
	//   description: name of user
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggUsersResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	m.Methods("GET").Path("/api/users/{username}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.GetUDFEndpoint, decodeGetUDFRequest, encodeGetUDFResponse, options...)),
	).Queries("page", "{page}", "offset", "{offset}", "limit", "{limit}", "orderby", "{orderby}")
}

// decodeGetUDFRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUDFRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name, ok := vars["username"]
	if !ok {
		return nil, errors.New("not a valid username")
	}
	err, page, offset, limit, orderby := getQueries(r)
	req := endpoint.GetUDFRequest{
		Name:    name,
		Page:    page,
		Offset:  offset,
		Limit:   limit,
		Orderby: orderby,
	}
	return req, err
}

// encodeGetUDFResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUDFResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostHandler creates the handler logic
func makePostHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:route POST /api/users users swaggCreateUserReq
	// Create a new user with the profile.
	// If the user's id is "exists", error will be returned.
	// responses:
	//   200: swaggNoReturnValue
	//   400: swaggBadReq
	m.Methods("POST").Path("/api/users").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"POST"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.PostEndpoint, decodePostRequest, encodePostResponse, options...)),
	)
}

// decodePostRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PostRequest{}
	err := json.NewDecoder(r.Body).Decode(&req.User)
	return req, err
}

// encodePostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePatchHandler creates the handler logic
func makePatchHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/patch").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PatchEndpoint, decodePatchRequest, encodePatchResponse, options...)))
}

// decodePatchRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePatchRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PatchRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePatchResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePatchResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePutHandler creates the handler logic
func makePutHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation PUT /api/users/{userid} users swaggPutReq
	// ---
	// summary: Update the user profile
	// description: Update the user profile with the profile. Also, you need to specify the user ID.
	// parameters:
	// - name: userid
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/User"
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	m.Methods("PUT").Path("/api/users/{userid:[0-9]+}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"PUT"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.PutEndpoint, decodePutRequest, encodePutResponse, options...)),
	)
}

// decodePutRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePutRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PutRequest{}
	err := json.NewDecoder(r.Body).Decode(&req.User)
	vars := mux.Vars(r)
	id, ok := vars["userid"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req.Id = id
	return req, err
}

// encodePutResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePutResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteHandler creates the handler logic
func makeDeleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	// swagger:operation DELETE /api/users/{userid} users swaggDeleteReq
	// ---
	// summary: Delete the user
	// description: You need to specify the user ID to delete the user .
	// parameters:
	// - name: userid
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	m.Methods("DELETE").Path("/api/users/{userid:[0-9]+}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(http.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...)),
	)
}

// decodeDeleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["userid"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
