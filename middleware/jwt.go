package middleware

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenName = "token"
const Issuer = "Money-Hub"
const SecretKey = "MoneyDodo"

/*
*	Role
* 	0 - admin
* 	1 - student
* 	2 - enterprise
 */

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Id   string `json:"id"`
	Role int    `json:"role"`
	Auth bool   `json:"auth"`
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(SecretKey []byte, issuer string, id string, role int, auth bool) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		id,
		role,
		auth,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

/**
 * 解析 token
 */
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}
