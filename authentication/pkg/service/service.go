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

// User验证Response
type UserRes struct {
	Openid string `json:"openid"`
	Token  string `json:"token"`
}

// AuthenticationService describes the service.
type AuthenticationService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	// Students-微信小程序
	GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data *UserRes)

	// Admin-Web
	AdminLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string)

	// 企业
	EnterpriseLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string)

	// 登出
	Logout(ctx context.Context) (status bool, errinfo string, data string)
}

type basicAuthenticationService struct {
	*db.DBService
}

func saveToken(b *basicAuthenticationService, token string, id string) (err error) {
	// 将token保存或者更新进数据库中
	item := &model.Token{
		Id: id,
	}
	if has, _ := b.Engine().Get(item); has == false {
		item.Token = token
		_, err = b.Engine().Insert(item)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		item.Token = token
		_, err = b.Engine().Where("id=?", item.Id).Update(item)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// 微信小程序用户获取OpenId
func (b *basicAuthenticationService) GetOpenid(ctx context.Context, code string) (status bool, errinfo string, data *UserRes) {
	// TODO implement the business logic of GetOpenid

	// https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	var AppID = "wx6f4b63c2710e1bae"
	var AppSecret = "2f11628e5f62c350247e8deb24a9814b"
	//构造url
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + AppID +
		"&secret=" + AppSecret +
		"&js_code=" + code +
		"&grant_type=authorization_code"

	res, err := http.Get(url)

	if err != nil {
		return false, err.Error(), nil
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
					fmt.Println(err.Error())
				}
			}
			token, _ := middleware.CreateToken([]byte(middleware.SecretKey), middleware.Issuer, info.Openid, 1, user.CertificationStatus)

			// 将token保存或者更新进数据库中
			err := saveToken(b, token, info.Openid)
			checkErr(err)

			data = &UserRes{
				Openid: info.Openid,
				Token:  token,
			}
			return true, "", data
		}
		return false, info.Errmsg, nil
	}
	return false, err.Error(), nil
}

// 验证Admin登陆
func (b *basicAuthenticationService) AdminLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string) {
	// TODO implement the business logic of AdminLogin
	// 判断是否已经记录
	fmt.Println(name, password)
	admin := &model.Admin{
		Name: name,
	}
	has, _ := b.Engine().Get(admin)
	if has == true && admin.Password == password {
		token, _ := middleware.CreateToken([]byte(middleware.SecretKey), middleware.Issuer, name, 0, 2)
		// 将token保存或者更新进数据库中
		err := saveToken(b, token, name)
		checkErr(err)
		return true, "", token
	} else if has == true && admin.Password != password {
		return false, "Password is incorrect", ""
	} else {
		return false, "No such a admin", ""
	}
}

func (b *basicAuthenticationService) EnterpriseLogin(ctx context.Context, name string, password string) (status bool, errinfo string, data string) {
	// TODO implement the business logic of EnterpriseLogin
	return status, errinfo, data
}

// 登出
func (b *basicAuthenticationService) Logout(ctx context.Context) (status bool, errinfo string, data string) {
	// TODO implement the business logic of Logout
	item := model.Token{
		Id: ctx.Value("id").(string),
	}
	_, err := b.Engine().Delete(item)
	if err != nil {
		checkErr(err)
		return false, "Exit Unsuccessfully", ""
	}
	return true, "Exit Successfully", ""
}

// NewBasicAuthenticationService returns a naive, stateless implementation of AuthenticationService.
func NewBasicAuthenticationService() AuthenticationService {
	basicAuthSvc := &basicAuthenticationService{
		&db.DBService{},
	}
	err := basicAuthSvc.Bind("conf/conf.moneydodo.yml")
	if err != nil {
		log.Printf("The AuthService failed to bind with mysql")
	}
	return basicAuthSvc
}

// New returns a AuthenticationService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthenticationService {
	var svc AuthenticationService = NewBasicAuthenticationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
