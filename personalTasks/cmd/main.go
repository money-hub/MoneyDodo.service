package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/personalTasks"
)

func NewServer() *negroni.Negroni {
	router := mux.NewRouter()
	initRoutes(router)
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

func initRoutes(router *mux.Router) {
	sub := router.PathPrefix("/api/users").Subrouter()
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	svc := personalTasks.NewBasicPTaskService()
	svc = personalTasks.LoggingMiddleware{logger, svc}

}
