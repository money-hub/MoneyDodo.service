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
	"log"
	"os"

	"github.com/codegangsta/negroni"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/cpt"
	_ "github.com/money-hub/MoneyDodo.service/swagger"
)

const defaultPort = "8005"

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	svc := cpt.NewBasicCptService()
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
	sub := router.PathPrefix("/api/tasks").Subrouter()
	// swagger:operation GET /api/tasks/{taskId} cpt swaggGetSpecReq
	// ---
	// summary: Get the specical task of the user (with id=userId).
	// description: Get the specical task. You need to specify the userId and taskId.
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
	// summary: Get all tasks of the user (with id=userId).
	// description: Get all task of the user. You need to specify the userId and taskId.
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTasksResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("GET").Path("").Handler(getAllHandler)
	// swagger:operation POST /api/tasks cpt swaggPostReq
	// ---
	// summary: Create a task.
	// description: Create a task. Also, you need to specify the userId and taskId.
	// parameters:
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/Task"
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggTaskResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("POST").Path("").Handler(postHandler)
	// swagger:operation PUT /api/tasks/{taskId}?action={action} cpt swaggPutReq
	// ---
	// summary: Update the task information
	// description: Update the task information. Also, you need to specify the user ID and task ID.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   type: string
	//   required: true
	// - name: action
	//   in: path
	//   description: action can be one of "release", "claim" and "finish"
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
	sub.Methods("PUT").Path("/{taskId:[0-9]+}?action={action}").Handler(putHandler)
	// swagger:operation DELETE /api/users/{userId}/tasks/{taskId}?status={status} cpt swaggDeleteReq
	// ---
	// summary: Delete\Cancel to Release\Cancel to Claim a task.
	// description: When status is equal to "non-released" or "finished", it means the creator want to delete it. When status is equal to "released", it means the creator want to cancel to release it. When status is equal to "claimed", it means the recipient want to cancel to claim it. Also, you need to specify the user ID and task ID.
	// parameters:
	// - name: taskId
	//   in: path
	//   description: id of task
	//   required: true
	// - name: state
	//   in: path
	//   description: state can be one of "non-released", "released", "claimed" and "finished"
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggNoReturnValue"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("DELETE").Path("/{taskId:[0-9]+}?state={state}").Handler(deleteHandler).Queries("status", "{status}")

	n.UseHandler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	n.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for cpt service", port)
}
