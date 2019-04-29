// Copyright 2019 money-hub. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// money-hub MoneyDodo/personalTasks
//
// This documentation describes example APIs found under https://github.com/ribice/golang-swaggerui-example
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
	"log"
	"os"

	"github.com/codegangsta/negroni"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/personalTasks"
	_ "github.com/money-hub/MoneyDodo.service/swagger"
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
	// swagger:operation GET /api/users/{userId}/tasks/{taskId} personalTasks swaggGetSpecReq
	// ---
	// summary: Get the specical task of the user (with id=userId).
	// description: Get the specical task. You need to specify the userId and taskId.
	// parameters:
	// - name: userId
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
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
	sub.Methods("GET").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(getSpecHandler)
	// swagger:operation GET /api/users/{userId}/tasks personalTasks swaggGetAllReq
	// ---
	// summary: Get all tasks of the user (with id=userId).
	// description: Get all task of the user. You need to specify the userId and taskId.
	// parameters:
	// - name: userId
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTasksResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("GET").Path("/{userId:[0-9]+}/tasks").Handler(getAllHandler)
	// swagger:operation POST /api/users/{userId}/tasks/{taskId} personalTasks swaggPostClaimReq
	// ---
	// summary: Claim the task (with id=taskId).
	// description: The user (with id=userId) claims the task (with id=taskId). You need to specify the userId and taskId.
	// parameters:
	// - name: userId
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("POST").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(postClaimHandler)
	// swagger:operation POST /api/users/{userId}/tasks personalTasks swaggPostReq
	// ---
	// summary: Create a task.
	// description: Create a task. Also, you need to specify the userId and taskId.
	// parameters:
	// - name: userId
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/Task"
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTaskResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("POST").Path("/{userId:[0-9]+}/tasks").Handler(postHandler)
	// swagger:operation PUT /api/users/{userId}/tasks/{taskId} personalTasks swaggPutReq
	// ---
	// summary: Update the task information
	// description: Update the task information. Also, you need to specify the user ID and task ID.
	// parameters:
	// - name: userId
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/Task"
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("PUT").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(putHandler)
	// swagger:operation DELETE /api/users/{userId}/tasks/{taskId}?status={status} personalTasks swaggDeleteReq
	// ---
	// summary: Delete\Cancel to Release\Cancel to Claim a task.
	// description: When status is equal to "non-released" or "finished", it means the creator want to delete it. When status is equal to "released", it means the creator want to cancel to release it. When status is equal to "claimed", it means the recipient want to cancel to claim it. Also, you need to specify the user ID and task ID.
	// parameters:
	// - name: userId
	//   in: path
	//   description: id of user
	//   type: string
	//   required: true
	// - name: taskId
	//   in: path
	//   description: id of task
	//   required: true
	// - name: status
	//   in: path
	//   description: status of task
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("DELETE").Path("/{userId:[0-9]+}/tasks/{taskId:[0-9]+}").Handler(deleteHandler).Queries("status", "{status}")

	n.UseHandler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	n.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for personalTasks service", port)
}
