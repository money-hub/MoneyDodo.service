package charge

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/model"
)

type Decodes struct {
	GetSpecDecode      func(ctx context.Context, r *http.Request) (interface{}, error)
	GetAllDecode       func(ctx context.Context, r *http.Request) (interface{}, error)
	PostDecode         func(ctx context.Context, r *http.Request) (interface{}, error)
	DeleteDecode       func(ctx context.Context, r *http.Request) (interface{}, error)
	GetAllOfUserDecode func(ctx context.Context, r *http.Request) (interface{}, error)
}

type Encodes struct {
	GetSpecEncode      func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	GetAllEncode       func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	PostEncode         func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	DeleteEncode       func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
	GetAllOfUserEncode func(ctx context.Context, w http.ResponseWriter, response interface{}) (err error)
}

func MakeServerDecodes() Decodes {
	return Decodes{
		GetSpecDecode:      decodeRequest,
		GetAllDecode:       decodeRequest,
		PostDecode:         decodeRequest,
		DeleteDecode:       decodeRequest,
		GetAllOfUserDecode: decodeRequest,
	}
}

func MakeServerEncodes() Encodes {
	return Encodes{
		GetSpecEncode:      encodeResponse,
		GetAllEncode:       encodeResponse,
		PostEncode:         encodeResponse,
		DeleteEncode:       encodeResponse,
		GetAllOfUserEncode: encodeResponse,
	}
}

// decodeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded requestfrom the HTTP request body.
func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	page, offset, limit := 1, 0, -1
	orderby := "+id"
	var err error
	// 解析Queries
	vals := r.URL.Query()
	vars := mux.Vars(r)
	userId := vars["userId"]
	offsets, ok := vals["offset"]
	if ok {
		offset, err = strconv.Atoi(offsets[0])
		if err != nil {
			return Request{}, err
		}
	}
	pages, ok := vals["page"]
	if ok {
		page, err = strconv.Atoi(pages[0])
		if err != nil {
			return Request{}, err
		}
	}
	limits, ok := vals["limit"]
	if ok {
		limit, err = strconv.Atoi(limits[0])
		if err != nil {
			return Request{}, nil
		}
	}
	orderbys, ok := vals["orderby"]
	if ok {
		orderby = orderbys[0]
	}

	//解析路径中显示定义的参数
	chargeId := vars["chargeId"]
	req := Request{
		ChargeId: chargeId,
		UserId:   userId,
		Page:     page,
		Offset:   offset,
		Limit:    limit,
		Orderby:  orderby,
	}

	if page <= 0 || offset < 0 || (orderby != "+id" && orderby != "-id") {
		return Request{}, errors.New("The url queries are not correct.")
	}

	if r.Method == "POST" {
		wrapper := model.Wrapper{}
		err := json.NewDecoder(r.Body).Decode(&wrapper)
		if err != nil {
			log.Println("The upload wrapper is not correct.")
		}
		c := model.Charge{}
		err = json.Unmarshal(wrapper.Raw, &c)
		if err != nil {
			log.Println("task is not in the request Body.")
		}
		req.Charge = c
	}
	// fmt.Printf("%#v", req)
	return req, nil
}

// encodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
