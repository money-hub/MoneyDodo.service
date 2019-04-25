package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/personalTasks"
)

const defaultPort = "9090"

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	svc := personalTasks.NewBasicPTaskService()
	svc = &personalTasks.LoggingMiddleware{Logger: logger, Next: svc}
	eps := personalTasks.MakeServerEndpoints(svc)
	decodes := personalTasks.MakeServerDecodes()
	encodes := personalTasks.MakeServerEncodes()

	getSpecHandler := httptransport.NewServer(
		eps.GetSpecEndpoint,
		decodes.GetSpecDecode,
		encodes.GetSpecEncode,
	)

	getAllHandler := httptransport.NewServer(
		eps.GetAllEndpoint,
		decodes.GetAllDecode,
		encodes.GetAllEncode,
	)

	postClaimHandler := httptransport.NewServer(
		eps.PostClaimEndpoint,
		decodes.PostClaimDecode,
		encodes.PostClaimEncode,
	)

	postHandler := httptransport.NewServer(
		eps.PostEndpoint,
		decodes.PostDecode,
		encodes.PostEncode,
	)

	putHandler := httptransport.NewServer(
		eps.PutEndpoint,
		decodes.PutDecode,
		encodes.PutEncode,
	)

	deleteHandler := httptransport.NewServer(
		eps.DeleteEndpoint,
		decodes.DeleteDecode,
		encodes.DeleteEncode,
	)

	n := negroni.Classic()
	router := mux.NewRouter()
	sub := router.PathPrefix("/api/users").Subrouter()
	sub.Methods("GET").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(getSpecHandler)
	sub.Methods("GET").Path("/{userId:[0-9]+}/tasks").Handler(getAllHandler)
	sub.Methods("POST").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(postClaimHandler)
	sub.Methods("POST").Path("/{userId:[0-9]+}/tasks").Handler(postHandler)
	sub.Methods("PUT").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(putHandler)
	sub.Methods("DELETE").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(deleteHandler).Queries("status", "{status}")

	n.UseHandler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	n.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for dailyhub server", port)
}
