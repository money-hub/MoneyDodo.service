# 创建相应的数据库
create database if not exists MoneyDodo;

use MoneyDodo;
drop table if exists enterprise;
drop table if exists admin;
drop table if exists deal;
drop table if exists comment;
drop table if exists questionnaire;
drop table if exists user;
drop table if exists task;
drop table if exists token;

# 用户
create table if not exists user (
	id varchar(100) not null primary key COMMENT 'OpenId',
    name varchar(100) COMMENT '姓名',
    sId varchar(20) COMMENT '学号',
    introduction text COMMENT '个人简介',
    balance double COMMENT '余额',
    icon mediumtext COMMENT '头像',
    phone varchar(11) COMMENT '电话号码',
    creditScore int COMMENT '信用分数',
    email varchar(50) COMMENT '邮箱',
    certificationStatus int DEFAULT 0 COMMENT '0-未提交，1-已提交未认证，2-审核通过，3-审核驳回',
    certifiedPic mediumtext COMMENT '认证图片'
);

# 管理员
create table if not exists admin (
    name varchar(20) not null primary key COMMENT '姓名',
    password varchar(20) not null COMMENT '密码'
);

# 企业
create table if not exists enterprise (
    name varchar(20) not null primary key COMMENT '姓名',
    password varchar(20) not null COMMENT '密码'
);

# 任务
create table if not exists task (
	id int AUTO_INCREMENT primary key COMMENT '任务id',
    kind varchar(20) not null COMMENT '任务类型',
    publisher varchar(20) not null COMMENT '发布者',
    -- recipient varchar(20) COMMENT '接收者',
    restrain text COMMENT '任务限制',
    pubdate text COMMENT '发布时间',
    cutoff text COMMENT '截至时间',
    -- enddate text COMMENT '结束时间',
    reward double COMMENT '赏金金额',
    -- recipientFinish bool COMMENT '接收者确认完成',
    -- ConfirmFinish bool COMMENT '接收者确认完成',
    state varchar(20) COMMENT '任务状态'
);

# 问卷
create table if not exists questionnaire (
    taskId int not null COMMENT '任务Id',
    query mediumtext COMMENT '填空',
    singleChoice mediumtext COMMENT '单项选择',
    primary key(taskId),
    foreign Key(taskId) references task(id)
);

# 交易
create table if not exists deal (
    id int AUTO_INCREMENT COMMENT '交易Id',
    taskId int not null COMMENT '任务Id',
    publisher varchar(20) COMMENT '发布者',
    recipient varchar(20) COMMENT '接受者',
    since text COMMENT '交易开始时间',
    until text COMMENT '交易结束时间',
    reward double COMMENT '交易额',
    state varchar(20) COMMENT '交易状态',
    primary key(id),
    foreign key(taskId) references task(id),
    foreign key(publisher) references user(id),
    foreign key(recipient) references user(id)
);

# 评论
create table if not exists comment (
    id int AUTO_INCREMENT COMMENT '评论Id',
    taskId int not null COMMENT '任务Id',
    userId varchar(20) not null COMMENT '用户Id',
    timestamp text COMMENT '评论时间戳',
    content mediumtext COMMENT '评论内容，支持图片评论',
    stars int default 0 COMMENT '评论点赞数量',
    stargazers text COMMENT '评论者的id'
    primary key(id),
    foreign key(taskId) references task(id),
    foreign key(userId) references user(id)
);

# Token
create table if not exists token(
	id char(255) COMMENT '唯一标识，包含管理员、企业、学生',
    token char(255),
    primary key(id)
);

create table if not exists recharge(
    id int AUTO_INCREMENT COMMENT '余额Id',
    amount int COMMENT '充值金额',
    timestamp text COMMENT '充值时间戳',
);