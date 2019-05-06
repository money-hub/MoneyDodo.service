package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type handle struct {
	host string
	port string
}

type Service struct {
	auth    *handle
	user    *handle
	task    *handle
	certify *handle
}

func (this *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// remote, err := url.Parse("http://" + this.host + ":" + this.port)
	// if err != nil {
	// 	panic(err)
	// }
	// proxy := httputil.NewSingleHostReverseProxy(remote)
	// proxy.ServeHTTP(w, r)
	fmt.Println(r.RequestURI)
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, token")
	// w.Header().Set("Access-Control-Allow-Methods", "*")
	// w.Header().Set("Content-Type", "application/json")
	var remote *url.URL
	if strings.Contains(r.RequestURI, "api/auth/") {
		remote, _ = url.Parse("http://" + this.auth.host + ":" + this.auth.port)
	} else if strings.Contains(r.RequestURI, "api/user/") {
		remote, _ = url.Parse("http://" + this.user.host + ":" + this.user.port)
	} else if strings.Contains(r.RequestURI, "api/users/") {
		remote, _ = url.Parse("http://" + this.certify.host + ":" + this.certify.port)
	} else {
		fmt.Fprintf(w, "404 Not Found")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	// 注册被代理的服务器 (host， port)
	service := &Service{
		auth:    &handle{host: "127.0.0.1", port: "8081"},
		user:    &handle{host: "127.0.0.1", port: "8081"},
		certify: &handle{host: "127.0.0.1", port: "8081"},
	}
	err := http.ListenAndServe(":8888", service)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
