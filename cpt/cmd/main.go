// Copyright 2019 money-hub. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// money-hub MoneyDodo/cpt
//
// This documentation describes example APIs found under https://github.com/money-hub/MoneyDodo.service
//
//     Schemes: http
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"flag"
	"log"
	"os"

	"github.com/codegangsta/negroni"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/cpt"
	"github.com/money-hub/MoneyDodo.service/middleware"
	_ "github.com/money-hub/MoneyDodo.service/swagger"
)

const defaultConf = "conf/conf.moneydodo.yml"
const defaultPort = "8005"

func main() {
	conf := flag.String("conf", defaultConf, "database config file")

	logger := kitlog.NewLogfmtLogger(os.Stderr)
	svc := cpt.NewBasicCptService(*conf)
	svc = &cpt.LoggingMiddleware{Logger: logger, Next: svc}
	eps := cpt.MakeServerEndpoints(svc)
	decodes := cpt.MakeServerDecodes()
	encodes := cpt.MakeServerEncodes()

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
	router.Use(middleware.GetTokenInfo)
	sub := router.PathPrefix("/api/tasks").Subrouter()
	// swagger:operation GET /api/tasks/{taskId} cpt swaggGetSpecReq
	// ---
	// summary: Get the specical task of the user (with id=userId).
	// description: Get the specical task with detail(Now only support questionnaire). You need to specify and taskId.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTaskResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("GET").Path("/{taskId:[0-9]+}").Handler(getSpecHandler)
	// swagger:operation GET /api/tasks cpt swaggGetAllReq
	// ---
	// summary: Get all tasks.
	// description: Get all task without detail.
	// parameters:
	// - name: page
	//   in: query
	//   description: page indicates the number of pages you want to get from server.
	//   type: int
	// - name: offset
	//   in: query
	//   description: offset indicates the number of targets you want to skip.
	//   type: int
	// - name: limit
	//   in: query
	//   description: limit indicates the number of targets in one page you want to get from server.
	//   type: int
	// - name: orderby
	//   in: query
	//   description: orderby indicates the order of targets you want to get from server.
	//   type: int
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTasksResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("GET").Path("").Handler(getAllHandler)
	// swagger:operation POST /api/tasks cpt swaggPostReq
	// ---
	// summary: Create a task.
	// description: Create a task(Now only support questionnaire). The upload parameter needs to be a wrapper of model.Qtnr and has one extra field called "kind" which indicates the task type.
	// parameters:
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/Qtnr"
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTaskResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("POST").Path("").Handler(postHandler)
	// swagger:operation PUT /api/tasks/{taskId} cpt swaggPutReq
	// ---
	// summary: Update the task information
	// description: Update the task information(Now only support questionnaire). Also, you need to specify the taskId. The upload parameter needs to be a wrapper of model.Qtnr and has one extra field called "kind" which indicates the task type.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/Qtnr"
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTaskResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("PUT").Path("/{taskId:[0-9]+}").Handler(putHandler)
	// swagger:operation DELETE /api/tasks/{taskId}?state={state} cpt swaggDeleteReq
	// ---
	// summary: Delete\Cancel a task.
	// description: When state is equal to "non-released" or "closed", it means the creator want to delete it. When state is equal to "released", it means the creator want to cancel to release it. Also, you need to specify the taskId.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   required: true
	// - name: state
	//   in: query
	//   description: state can be one of "non-released", "released" and "closed"
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("DELETE").Path("/{taskId:[0-9]+}").Handler(deleteHandler).Queries("state", "{state}")

	n.UseHandler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	n.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for cpt service", port)
}
