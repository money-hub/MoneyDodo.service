// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	endpoint "github.com/money-hub/MoneyDodo.service/task/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeGetHisReleasedTasksHandler(m, endpoints, options["GetHisReleasedTasks"])
	makeGetTasksByIDHandler(m, endpoints, options["GetTasksByID"])
	makeGetHisUnreleasedTasksHandler(m, endpoints, options["GetHisUnreleasedTasks"])
	makeGetHisClosedTasksHandler(m, endpoints, options["GetHisClosedTasks"])
	return m
}
