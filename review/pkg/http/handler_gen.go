// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	middleware "github.com/money-hub/MoneyDodo.service/middleware"
	endpoint "github.com/money-hub/MoneyDodo.service/review/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	m.Use(middleware.GetTokenInfo)
	makePostReviewHandler(m, endpoints, options["PostReview"])
	makeGetReviewHandler(m, endpoints, options["GetReview"])
	makeGetReviewsHandler(m, endpoints, options["GetReviews"])
	makePutReviewHandler(m, endpoints, options["PutReview"])
	return m
}
