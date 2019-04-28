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
    recipient text COMMENT '接收者',
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
