package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/money-hub/MoneyDodo.service/db"
	MyJwt "github.com/money-hub/MoneyDodo.service/middleware"
	"github.com/money-hub/MoneyDodo.service/model"
)

type handle struct {
	host string
	port string
}

type Service struct {
	auth     *handle
	certify  *handle
	user     *handle
	task     *handle
	cpt      *handle
	comment  *handle
	deal     *handle
	txn      *handle
	balance  *handle
	recharge *handle
}

type RespData struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    interface{} `json:"data"`
}

/**
*	过滤规则
*	Admin - 全部请求都接受
*	User - 未实名验证则权限受限
**/

// User
var PublicURL = []string{
	// 登陆服务（authentication） get/post
	"/api/auth",
	// 查询某个用户发布的任务 get
	"/api/users/[a-zA-Z0-9]+/tasks?state=released",
	// 查询某个任务 get
	"/api/tasks[/0-9]*",
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getBasicService(conf string) *db.DBService {
	basicSvc := &db.DBService{}
	err := basicSvc.Bind(conf)
	if err != nil {
		log.Printf("The Proxy Service failed to bind with mysql")
	}
	return basicSvc
}

func writeResp(status bool, errinfo string, data interface{}) []byte {
	RespData := RespData{
		Status:  status,
		Errinfo: errinfo,
		Data:    data,
	}
	respose, err := json.Marshal(RespData)
	if err != nil {
		log.Fatalln(err)
	}
	return respose
}

func (this *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RequestURI)
	basicSvc := getBasicService("conf/conf.lyt.yml")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, token")
	// w.Header().Set("Access-Control-Allow-Methods", "*")
	// w.Header().Set("Content-Type", "application/json")
	var remote *url.URL

	if strings.HasPrefix(r.RequestURI, "/api/auth") && r.RequestURI != "/api/auth/logout" {
		// 登陆服务（authentication）
		remote, _ = url.Parse("http://" + this.auth.host + ":" + this.auth.port)
	} else if match, _ := regexp.MatchString("/api/users/[a-zA-Z0-9]+/tasks?state=released", r.RequestURI); match && strings.ToUpper(r.Method) == "GET" {
		// 用户任务信息（task） - 查询某个用户发布的任务 get
		remote, _ = url.Parse("http://" + this.task.host + ":" + this.task.port)
	} else if match, _ := regexp.MatchString("/api/tasks[/0-9]*", r.RequestURI); match && strings.ToUpper(r.Method) == "GET" {
		// 用户任务信息（task） - 查询任务 get
		remote, _ = url.Parse("http://" + this.cpt.host + ":" + this.cpt.port)
	} else {
		var mapClaims jwt.MapClaims
		myToken := ""
		// 如果token存在于Authorization中
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte(MyJwt.SecretKey), nil
		})
		if token != nil {
			var ok bool
			mapClaims, ok = token.Claims.(jwt.MapClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(writeResp(false, "Unauthorized access to this resource", nil))
				return
			}
			myToken = strings.Split(r.Header["Authorization"][0], " ")[1]
		} else {
			// 如果token存在于header中
			for k, v := range r.Header {
				if strings.ToLower(k) == "token" {
					myToken = v[0]
					break
				}
			}
			mapClaims, err = MyJwt.ParseToken(myToken, []byte(MyJwt.SecretKey))
			checkErr(err)
			if err != nil || mapClaims.Valid() != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(writeResp(false, "Unauthorized access to this resource", nil))
				return
			}
		}

		if myToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(writeResp(false, "Unauthorized access to this resource", nil))
			return
		}

		// 校验token是否合法
		item := &model.Token{
			Id: mapClaims["id"].(string),
		}

		has, _ := basicSvc.Engine().Get(item)
		// 判断是否为认证信息相关
		match, _ := regexp.MatchString("/api/users/[a-zA-Z0-9]+/certs", r.RequestURI)
		if has == false || item.Token != myToken || (int(mapClaims["role"].(float64)) != 0 && !match && int(mapClaims["certificationStatus"].(float64)) != 2) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(writeResp(false, "Unauthorized access to this resource", nil))
			return
		}

		// 信息放入上下文中
		// ctx := context.WithValue(r.Context(), "id", mapClaims["id"])
		// ctx = context.WithValue(ctx, "role", mapClaims["role"])
		// ctx = context.WithValue(ctx, "certificationStatus", mapClaims["certificationStatus"])
		// r = r.WithContext(ctx)

		// 登陆服务（authentication） 登出
		if r.RequestURI == "/api/auth/logout" {
			remote, _ = url.Parse("http://" + this.auth.host + ":" + this.auth.port)
		}

		// 实名认证（certify）
		if match, _ := regexp.MatchString("/api/users/[a-zA-Z0-9]+/certs", r.RequestURI); match || strings.HasPrefix(r.RequestURI, "/api/certs") {
			remote, _ = url.Parse("http://" + this.certify.host + ":" + this.certify.port)
		}

		// user相关 - /api/users
		if strings.HasPrefix(r.RequestURI, "/api/users") {
			if match, _ := regexp.MatchString("/api/users/[a-zA-Z0-9]+/tasks", r.RequestURI); match || strings.HasPrefix(r.RequestURI, "/api/certs") {
				// 用户任务信息（task）
				remote, _ = url.Parse("http://" + this.task.host + ":" + this.task.port)
			} else if match, _ := regexp.MatchString("/api/users/[a-zA-Z0-9]+/deals", r.RequestURI); match || strings.HasPrefix(r.RequestURI, "/api/certs") {
				// 用户交易信息（deal）
				remote, _ = url.Parse("http://" + this.deal.host + ":" + this.deal.port)
			} else if match, _ := regexp.MatchString("/api/users/[a-zA-Z0-9]+/balances", r.RequestURI); match || strings.HasPrefix(r.RequestURI, "/api/certs") {
				// 充值信息（balance）
				remote, _ = url.Parse("http://" + this.balance.host + ":" + this.balance.port)
			} else {
				// 个人信息（user）
				remote, _ = url.Parse("http://" + this.user.host + ":" + this.user.port)
			}
		}

		// taks相关 - /api/tasks
		if strings.HasPrefix(r.RequestURI, "/api/tasks") {
			if match, _ := regexp.MatchString("/api/tasks/[0-9]+/comments", r.RequestURI); match || strings.HasPrefix(r.RequestURI, "/api/certs") {
				// 任务评论（comment）
				remote, _ = url.Parse("http://" + this.comment.host + ":" + this.comment.port)
			} else {
				// 任务交互（cpt）
				remote, _ = url.Parse("http://" + this.cpt.host + ":" + this.cpt.port)
			}
		}

		// 交易操作（txn）
		if strings.HasPrefix(r.RequestURI, "/api/deals") {
			remote, _ = url.Parse("http://" + this.txn.host + ":" + this.txn.port)
		}

		// 充值（recharge）
		if strings.HasPrefix(r.RequestURI, "/api/balances") {
			remote, _ = url.Parse("http://" + this.recharge.host + ":" + this.recharge.port)
		}
	}
	if remote == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(writeResp(false, "404 Not Found", nil))
	}
	// 代理路由分发
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	// 注册被代理的服务器 (host， port)
	service := &Service{
		auth:     &handle{host: "127.0.0.1", port: "8001"},
		certify:  &handle{host: "127.0.0.1", port: "8002"},
		user:     &handle{host: "127.0.0.1", port: "8003"},
		task:     &handle{host: "127.0.0.1", port: "8004"},
		cpt:      &handle{host: "127.0.0.1", port: "8005"},
		comment:  &handle{host: "127.0.0.1", port: "8006"},
		deal:     &handle{host: "127.0.0.1", port: "8007"},
		txn:      &handle{host: "127.0.0.1", port: "8008"},
		balance:  &handle{host: "127.0.0.1", port: "8009"},
		recharge: &handle{host: "127.0.0.1", port: "8010"},
	}
	err := http.ListenAndServe(":8888", service)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
