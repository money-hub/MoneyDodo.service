package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type WeChatRes struct {
	Openid      string `json:"openid"`
	Session_key string `json:"session_key"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

// AuthenticationService describes the service.
type AuthenticationService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	// data - token

	// 微信小程序
	GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data string)

	// 管理员Web
	AdminLogin(ctx context.Context) (status bool, errinfo string, data string)
}

type basicAuthenticationService struct{}

func (b *basicAuthenticationService) GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data string) {
	// TODO implement the business logic of GetOpenid

	// https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code

	var AppId = "wx6f4b63c2710e1bae"
	var AppSecret = "2f11628e5f62c350247e8deb24a9814b"

	//构造url
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + AppId +
		"&secret=" + AppSecret +
		"&js_code=" + code +
		"&grant_type=authorization_code"

	res, err := http.Get(url)

	if err != nil {
		return false, err.Error(), ""
	}

	defer res.Body.Close() //关闭链接

	body, _ := ioutil.ReadAll(res.Body)
	var info WeChatRes
	if err := json.Unmarshal(body, &info); err == nil {
		if info.Errcode == 0 {
			return true, "", info.Openid
		} else {
			return false, info.Errmsg, ""
		}
	} else {
		return false, err.Error(), ""
	}
}

func (b *basicAuthenticationService) AdminLogin(ctx context.Context) (status bool, errinfo string, data string) {
	// TODO implement the business logic of AdminLogin

	return status, errinfo, data
}

// NewBasicAuthenticationService returns a naive, stateless implementation of AuthenticationService.
func NewBasicAuthenticationService() AuthenticationService {
	return &basicAuthenticationService{}
}

// New returns a AuthenticationService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthenticationService {
	var svc AuthenticationService = NewBasicAuthenticationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
