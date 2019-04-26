# 创建相应的数据库
create database if not exists MoneyDodo;

use MoneyDodo;

# 用户
create table if not exists user (
	id varchar(20) not null primary key COMMENT 'OpenId',
    name varchar(20) COMMENT '姓名',
    sId varchar(20) COMMENT '学号',
    introduction text COMMENT '个人简介',
    balance double COMMENT '余额',
    icon MEDIUMBLOB COMMENT '头像',
    phone varchar(11) COMMENT '电话号码',
    creditScore int COMMENT '信用分数',
    email varchar(20) COMMENT '邮箱',
    certificationStatus int DEFAULT 0 COMMENT '0-未提交，1-已提交未认证，2-审核通过，3-审核驳回',
    certifiedPic MEDIUMBLOB COMMENT '认证图片'
);

# 管理员
create table if not exists admin (
    name varchar(20) not null COMMENT '姓名',
    password varchar(20) not null COMMENT '密码'
);

# 企业
create table if not exists enterprise (
    name varchar(20) not null COMMENT '姓名',
    password varchar(20) not null COMMENT '密码'
);

# 任务
create table if not exists task (
	id int AUTO_INCREMENT primary key COMMENT '任务id',
    type varchar(20) not null COMMENT '任务类型',
    publisher varchar(20) not null COMMENT '发布者',
    recipient text COMMENT '接收者',package service

import (
	"context"
	"log"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// CertifyService describes the service.
type CertifyService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetAllUnAuth(ctx context.Context) (status bool, errinfo string, data []model.User)
	PostAuthInfo(ctx context.Context, id string, img []byte) (status bool, errinfo string)
	PostCertifyInfo(ctx context.Context, id string, pass bool) (status bool, errinfo string)
}

type basicCertifyService struct {
	*db.DBService
}

func (b *basicCertifyService) GetAllUnAuth(ctx context.Context) (status bool, errinfo string, data []model.User) {
	// TODO implement the business logic of GetAllUnAuth
	user := model.User{}
	rows, err := b.Engine().Where("certificationStatus = ?", 1).Rows(user)
	if err == nil {
		for rows.Next() {
			err1 := rows.Scan(user)
			if err1 != nil {
				return false, err1.Error(), data
			}
			data = append(data, user)
		}
		return true, "", data
	} else {
		return false, err.Error(), data
	}
}
func (b *basicCertifyService) PostAuthInfo(ctx context.Context, id string, img []byte) (status bool, errinfo string) {
	// TODO implement the business logic of PostAuthInfo
	user := model.User{
		Id: id,
	}
	status, err := b.Engine().Get(user)
	if status == false || err != nil {
		return false, err.Error()
	}
	user.CertifiedPic = img
	user.CertificationStatus = 1
	_, err = b.Engine().Id(id).Update(user)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}
func (b *basicCertifyService) PostCertifyInfo(ctx context.Context, id string, pass bool) (status bool, errinfo string) {
	// TODO implement the business logic of PostCertifyInfo
	user := model.User{
		Id: id,
	}
	status, err := b.Engine().Get(user)
	if status == false || err != nil {
		return false, err.Error()
	}
	if pass {
		user.CertificationStatus = 2
	} else {
		user.CertificationStatus = 3
	}
	_, err = b.Engine().Id(id).Update(user)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

// NewBasicCertifyService returns a naive, stateless implementation of CertifyService.
func NewBasicCertifyService() CertifyService {
	basicCertifyService := &basicCertifyService{
		&db.DBService{},
	}
	err := basicCertifyService.Bind("conf/conf.users.yml")
	if err != nil {
		log.Printf("The UserService failed to bind with mysql")
	}
	return basicCertifyService
}

// New returns a CertifyService with all of the expected middleware wired in.
func New(middleware []Middleware) CertifyService {
	var svc CertifyService = NewBasicCertifyService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

    restrain text COMMENT '任务限制',
    pubdate text COMMENT '发布时间',
    cutoff text COMMENT '截至时间',
    reward double COMMENT '赏金金额',
    status varchar(20) COMMENT '任务状态'
);

# 用户-任务
create table if not exists relation (
    userId varchar(20) not null COMMENT '用户Id',
    taskId int not null COMMENT '任务Id',
    detail varchar(20) not null COMMENT '发布或者接受',
    primary key(userId, taskId, detail),
    foreign key(userId) references user(id),
    foreign Key(taskId) references task(id)
);
