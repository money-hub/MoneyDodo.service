# MoneyDodo.service

[![Build Status](https://travis-ci.org/money-hub/MoneyDodo.service.svg?branch=master)](https://travis-ci.org/money-hub/MoneyDodo.service.svg?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/money-hub/MoneyDodo.service/badge.svg?branch=master)](https://coveralls.io/github/money-hub/MoneyDodo.service?branch=master)
[![Analytics](https://ga-beacon.appspot.com/UA-139167220-1/welcome-page)](https://github.com/money-hub/MoneyDodo.service)

此项目为MoneyDodo.web和MoneyDodo.wechat的后台服务实现，采用微服务，使用go-kit框架，代码主体利用[GoKit CLI](<https://github.com/kujtimiihoxha/kit>)开源工具快速生成，便于将开发重心放到业务逻辑上。

## 代码结构介绍

**conf**

存放数据库配置文件，具体配置方法参考`conf.example.yml`

**db**

存放数据库相关内容，初始化xorm、MySQL引擎

**model**

存放数据结构，包括user、task等

**middleware**

中间件，目前包括jwt认证中间件。

**swagger**

swaggerui，API的可视化界面，便于前后端进行交互

**authentication**

用户登录认证相关服务，主要处理用户登陆的认证，token获取，退出等请求

**user**

用户系统微服务，主要处理用户信息的Get、Post、Put、Patch、Delete等请求

**personalTasks**

用户task相关微服务，主要处理用户发布任务，删除任务，领取任务，查询任务等请求

**certify**

用户实名认知相关服务，主要处理用户上传实名认证信息，查询信息等请求。

**...其他微服务，待完成**

之后，代码结构可以再次调整，将共用代码进行提取，减少代码冗余。

## GoKit CLI使用方法

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

