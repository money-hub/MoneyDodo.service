#源镜像
FROM golang:1.12
#设置镜像工作目录
WORKDIR $GOPATH/src/github.com/money-hub/MoneyDodo.service
#将宿主机的go工程代码加入到docker容器中
ADD ../ $GOPATH/src/github.com/money-hub/MoneyDodo.service
# 安装依赖包
RUN go get ./...
# 设置 PORT 环境变量
# ENV PORT 8003
#暴露端口
EXPOSE 8003
#最终运行docker的命令
ENTRYPOINT ["go run", "./user/cmd/main.go"]