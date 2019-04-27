# MoneyDodo.service

[![Build Status](https://travis-ci.org/money-hub/MoneyDodo.service.svg?branch=master)](https://travis-ci.org/money-hub/MoneyDodo.service.svg?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/money-hub/MoneyDodo.service/badge.svg?branch=master)](https://coveralls.io/github/money-hub/MoneyDodo.service?branch=master)
[![Analytics](https://ga-beacon.appspot.com/UA-139167220-1/welcome-page)](https://github.com/money-hub/MoneyDodo.service)

此项目为MoneyDodo.web和MoneyDodo.wechat的后台服务实现，采用微服务，使用go-kit框架，代码主体利用[GoKit CLI](<https://github.com/kujtimiihoxha/kit>)开源工具快速生成，便于将开发重心放到业务逻辑上。

## 一、代码结构介绍

**1. conf:** 存放数据库配置文件，具体配置方法参考`conf.example.yml`

**2. db:** 存放数据库相关内容，初始化xorm、MySQL引擎

**3. model:** 存放数据结构，包括user、task等

**4. middleware:** 中间件，目前包括jwt认证中间件。

**5. swagger:** swaggerui，API的可视化界面，便于前后端进行交互

**6. authentication:** 用户登录认证相关服务，主要处理用户登陆的认证，token获取，退出等请求

**7. user:** 用户系统微服务，主要处理用户信息的Get、Post、Put、Patch、Delete等请求

**8. personalTasks:** 用户task相关微服务，主要处理用户发布任务，删除任务，领取任务，查询任务等请求

**9. certify:** 用户实名认知相关服务，主要处理用户上传实名认证信息，查询信息等请求。

**...其他微服务，待完成**

之后，代码结构可以再次调整，将共用代码进行提取，减少代码冗余。

## 二、GoKit CLI使用方法

**GoKit CLI**使用方法可以参考<https://medium.com/@kujtimii.h/creating-a-todo-app-using-gokit-cli-20f066a58e1>

```bash
# 以用户系统微服务为例
$ kit n s user
# 之后定义相关服务（Get、Post等）
$ kit g s user -w --gorilla
#-w generate some default service middleware.
#--gorilla use gorilla/mux instead of the default http handler for the http transport.
$ go run user/cmd/main.go
#ts=2019-04-16T09:52:56.5674053Z caller=service.go:78 tracer=none
#ts=2019-04-16T09:52:56.5714055Z caller=service.go:100 transport=HTTP addr=:8081
#ts=2019-04-16T09:52:56.5714055Z caller=service.go:134 transport=debug/HTTP addr=:8080
#端口为8081
```

## 三、swaggerui使用方法

参考链接：
https://swagger.io/docs/specification/paths-and-operations/
https://studygolang.com/articles/12354?fr=sidebar

### 以user服务中PUT请求为例：

#### A.注释规范

```bash
// swagger:operation PUT /api/users/{userid} users swaggPutReq
// ---
// summary: Update the user profile
// description: Update the user profile with the profile. Also, you need to specify the user ID.
// parameters:
// - name: userid
//   in: path
//   description: id of user
//   type: string
//   required: true
// - name: Body
//   in: body
//   schema:
//     "$ref": "#/definitions/User"
//   required: true
// responses:
//   "200":
//	   "$ref": "#/responses/swaggNoReturnValue"
//   "400":
//	   "$ref": "#/responses/swaggBadReq"
```

**1. swagger:operation** - 为提示符，表示一个请求操作

**2. PUT** - `HTTP`方法

**3. /api/users/{userid}** - 路径

**4. users** - 类似于路由分隔标签，将相同的分隔标签的请求归到同一组

**5. swaggPutReq** - 此参数没有具体意义，单参数是强制性的，但是推荐不同请求使用不同的参数。命名格式可采用**swaggXXXReq**，若不同请求该参数一样，会出现很多bug。

**6. ---** - 分隔符，下方代码为`YAML`格式的`swagger`规范，缩进必须保持一致且正确，推荐使用两格缩进。否则将无法正常解析。

**7. summary** - 标题，`API`的概括描述

**8. description** - 描述，`API`的详细描述

**9. parameters** - `URL`参数，此例子中为`{userId}`，如果需要`query`的，可使用`？name={name}`来表示

**10.  - name** - 指定参数，此例子中为`URL`中的`userId`

**11. in** - 表示此参数位于哪个部分，`path`表示位于`URL`路径中，`body`表示位于上传的`request body`中

**12. description** - 参数说明

**13. type** - 指定参数类型

**14. required** - 是否一定需要此参数

**15. schema** - 当参数位于`body`中需要此参数，指定参数的数据结构，**"$ref": "#/definitions/XXX"**按照此种格式写即可，具体原因尚未搞懂。

**16. responses** - 说明返回类型。

**17. "200"** - `200`表示状态码，我用`200`表示成功的请求；**"$ref": "#/responses/swaggNoReturnValue"**按照此格式来书写，其中`swaggNoReturnValue`定义在**swagger/model.go**文件中：

```bash
// HTTP status code 200 and no return value
// swagger:response swaggNoReturnValue
type swaggNoReturnValue struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Status bool `json:"status"`
		// Detailed error message
		Errinfo string `json:"errinfo"`
	}
}
```

- 第一行注释：尽量书写，会体现在`swaggerui`中
- 第二行指数：`swagger:response`为提示符，`swaggNoReturnValue`为下方数据类型的一个`tag`，这两者共同组成了**"$ref": "#/responses/swaggNoReturnValue"**
- 数据结构中，有三个属性：`Status`、`Errinfo`、`Data`，上述结构中没有返回值，所以取消了最后一个参数。

**18. "400"** - **400**表示状态码，我用`400`来表示失败的请求。返回格式说明与上述过程一致。

更加详细的说明参看[官方文档](https://swagger.io/docs/specification/paths-and-operations/)。

#### B. 生成swagger.user.json

- 注意在`user/cmd/main.go`文件中`import _ "github.com/money-hub/MoneyDodo.service/swagger"`，只有这样，上述注释中的`#/responses/`才会被识别。

- 安装**go-swagger**:

  ```bash
  go get github.com\go-swagger\go-swagger\cmd\swagger
  ```

- 生成spec

  ```bash
  # 当前路径：根目录
  $ cd user/cmd
  # 此命令会从当前文件及下的main函数入口递归搜索所有文件的swagger注释，最终生成指定的`swagger.users.json`
  $ swagger generate spec -o ../../swagger/swaggerui/dist/swagger.user.json
  $ cd ../..
  # 开启的默认端口为8000，注意不要占用此端口
  $ go run swagger/swaggerui/main.go
  ```

#### C. 说明

每个服务各自生成相应的文档，命名格式为`swagger.XXX.json`。`XXX`为服务名称，方便查询。