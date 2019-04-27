package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/middleware"
	"github.com/money-hub/MoneyDodo.service/model"
)

// WeChatRes 微信服务器Response
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

	// Students-微信小程序
	GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data string)

	// Admin-Web
	AdminLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string)
}

type basicAuthenticationService struct {
	*db.DBService
}

// 微信小程序用户获取OpenId
func (b *basicAuthenticationService) GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data string) {
	// TODO implement the business logic of GetOpenid

	// https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	// var AppID = "wx6f4b63c2710e1bae"
	// var AppSecret = "2f11628e5f62c350247e8deb24a9814b"
	var AppID = "wx25915d3c4f6a78f3"
	var AppSecret = "133e74afeca06c60a597cf3b694a6c87"

	//构造url
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + AppID +
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
			user := &model.User{
				Id: info.Openid,
			}
			// 判断是否已经记录
			has, _ := b.Engine().Get(user)
			if has == false {
				_, err := b.Engine().Insert(user)
				if err != nil {
					fmt.Println("[Authentication log] Insert user failed")
				}
			}
			token, _ := middleware.CreateToken([]byte(middleware.SecretKey), middleware.Issuer, info.Openid, 1, user.CertificationStatus)
			return true, "", token
		}
		return false, info.Errmsg, ""
	}
	return false, err.Error(), ""
}

// 验证Admin登陆
func (b *basicAuthenticationService) AdminLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string) {
	// TODO implement the business logic of AdminLogin
	// 判断是否已经记录
	admin := &model.Admin{
		Name: name,
	}
	has, _ := b.Engine().Get(admin)
	if has == true && admin.Password == password {
		token, _ := middleware.CreateToken([]byte(middleware.SecretKey), middleware.Issuer, name, 0, 0)
		return true, "", token
	} else if has == true && admin.Password != password {
		return false, "Password is incorrect", ""
	} else {
		return false, "No such a admin", ""
	}
}

// NewBasicAuthenticationService returns a naive, stateless implementation of AuthenticationService.
func NewBasicAuthenticationService() AuthenticationService {
	basicUserSvc := &basicAuthenticationService{
		&db.DBService{},
	}
	err := basicUserSvc.Bind("conf/conf.lyt.yml")
	if err != nil {
		log.Printf("The UserService failed to bind with mysql")
	}
	return basicUserSvc
}

// New returns a AuthenticationService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthenticationService {
	var svc AuthenticationService = NewBasicAuthenticationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
