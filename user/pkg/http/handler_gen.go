// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/middleware"
	endpoint "github.com/money-hub/MoneyDodo.service/user/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	m.Use(middleware.GetTokenInfo)
	makeGetSpecHandler(m, endpoints, options["GetSpec"])
	makeGetAllHandler(m, endpoints, options["GetAll"])
	makeGetUDFHandler(m, endpoints, options["GetUDF"])
	makePostHandler(m, endpoints, options["Post"])
	makePatchHandler(m, endpoints, options["Patch"])
	makePutHandler(m, endpoints, options["Put"])
	makeDeleteHandler(m, endpoints, options["Delete"])
	return m
}
