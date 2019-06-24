# JWT的使用

修改`service/pkg/http/handler_gen.go`的`NewHTTPHandler`函数，加入中间键处理
```go
// authentication/pkg/http/handler_gen.go

import MyJwt "github.com/money-hub/MoneyDodo.service/middleware"

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	m.Use(MyJwt.GetTokenInfo)   // 添加中间键处理
	makeGetOpenidHandler(m, endpoints, options["GetOpenid"])
	makeAdminLoginHandler(m, endpoints, options["AdminLogin"])
	makeEnterpriseLoginHandler(m, endpoints, options["EnterpriseLogin"])
	makeLogoutHandler(m, endpoints, options["Logout"])
	return m
}

```

```go
// 使用
eg:
ctx.Value("id").(string)
ctx.Value("role").(int)
```