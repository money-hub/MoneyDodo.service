// Copyright 2019 money-hub. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// money-hub MoneyDodo/charge
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
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/time/rate"

	"github.com/codegangsta/negroni"
	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/money-hub/MoneyDodo.service/charge"
	"github.com/money-hub/MoneyDodo.service/middleware"
	_ "github.com/money-hub/MoneyDodo.service/swagger"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const defaultConf = "conf/conf.moneydodo.yml"
const defaultPort = "8008"

func main() {
	conf := flag.String("conf", defaultConf, "database config file")

	logger := kitlog.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "moneydodo",
		Subsystem: "charge_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "moneydodo",
		Subsystem: "charge_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	limiter := rate.NewLimiter(rate.Every(time.Second*1), 1)

	svc := charge.NewBasicChargeService(*conf)
	svc = &charge.LoggingMiddleware{Logger: logger, Next: svc}
	svc = &charge.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}
	eps := charge.MakeServerEndpoints(svc)
	decodes := charge.MakeServerDecodes()
	encodes := charge.MakeServerEncodes()

	getAllOfUserHandler := httptransport.NewServer(
		ratelimit.NewDelayingLimiter(limiter)(eps.GetAllOfUserEndpoint),
		decodes.GetAllOfUserDecode,
		encodes.GetAllOfUserEncode,
	)

	getSpecHandler := httptransport.NewServer(
		ratelimit.NewDelayingLimiter(limiter)(eps.GetSpecEndpoint),
		decodes.GetSpecDecode,
		encodes.GetSpecEncode,
	)

	getAllHandler := httptransport.NewServer(
		ratelimit.NewDelayingLimiter(limiter)(eps.GetAllEndpoint),
		decodes.GetAllDecode,
		encodes.GetAllEncode,
	)

	postHandler := httptransport.NewServer(
		ratelimit.NewDelayingLimiter(limiter)(eps.PostEndpoint),
		decodes.PostDecode,
		encodes.PostEncode,
	)

	deleteHandler := httptransport.NewServer(
		ratelimit.NewDelayingLimiter(limiter)(eps.DeleteEndpoint),
		decodes.DeleteDecode,
		encodes.DeleteEncode,
	)

	n := negroni.Classic()
	router := mux.NewRouter()
	router.Use(middleware.GetTokenInfo)
	// swagger:operation GET /api/users/{userId}/charges charge swaggGetAllOfUserReq
	// ---
	// summary: Get all charges of user.
	// description: Get all charges of user.
	// parameters:
	// - name: userId
	//   in: path
	//   description: id of user
	//   type: string
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
	//	   "$ref": "#/responses/swaggChargesResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	router.Methods("GET").Path("/api/users/{userId}/charges").Handler(getAllOfUserHandler)

	sub := router.PathPrefix("/api/charges").Subrouter()
	sub.Path("/metrics").Handler(promhttp.Handler())
	// swagger:operation GET /api/charges/{chargeId} charge swaggGetSpecReq
	// ---
	// summary: Get the specical charge of the user (with id=userId).
	// description: Get the specical charge and you need to specify and chargeId.
	// parameters:
	// - name: chargeId
	//   in: path
	//   description: id of charge
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggChargeResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("GET").Path("/{chargeId:[0-9]+}").Handler(getSpecHandler)
	// swagger:operation GET /api/charges charge swaggGetAllReq
	// ---
	// summary: Get all charges.
	// description: Get all charges.
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
	//	   "$ref": "#/responses/swaggChargesResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("GET").Path("").Handler(getAllHandler)
	// swagger:operation POST /api/charges charge swaggPostReq
	// ---
	// summary: Create a charge.
	// description: Create a charge and the upload parameter needs to be model.Charge.
	// parameters:
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/Charge"
	// responses:
	//   "200":
	//	   "$ref": "#/responses/swaggChargeResp"
	//   "400":
	//	   "$ref": "#/responses/swaggBadReq"
	sub.Methods("POST").Path("").Handler(postHandler)
	// swagger:operation DELETE /api/charges/{chargeId} charge swaggDeleteReq
	// ---
	// summary: Delete a charge.
	// description: Delete a charge and you need to specify and chargeId.
	// parameters:
	// - name: chargeId
	//   in: path
	//   description: id of charge
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

	log.Printf("connect to http://localhost:%s/ for charge service", port)
}
