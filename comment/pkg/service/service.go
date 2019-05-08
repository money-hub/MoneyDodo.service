package service

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// CommentService describes the service.
type CommentService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	// 获取某个task的评论
	GetComment(ctx context.Context, taskId string) (status bool, errinfo string, data []model.Comment)

	// 发表评论
	PostComment(ctx context.Context, taskId string, comment string) (status bool, errinfo string, data *model.Comment)

	// 更改某条评论
	ChangeComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data *model.Comment)

	// 删除某条评论
	DeleteComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data string)

	// 点赞某条评论
	LikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment)

	// 取消点赞某条评论
	CancelLikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment)
}

type basicCommentService struct {
	*db.DBService
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// 获取某个task的评论
func (b *basicCommentService) GetComment(ctx context.Context, taskId string) (status bool, errinfo string, data []model.Comment) {
	// TODO implement the business logic of GetComment

	comments := make([]model.Comment, 0)
	err := b.Engine().Where("taskId = ?", taskId).Find(&comments)
	if err != nil {
		log.Println(err)
		return false, err.Error(), nil
	} else {
		return true, "", comments
	}
}

// 发表评论
func (b *basicCommentService) PostComment(ctx context.Context, taskId string, comment string) (status bool, errinfo string, data *model.Comment) {
	// TODO implement the business logic of PostComment

	// 判断对应任务是否存在
	task := &model.Task{
		Id: taskId,
	}
	ok, err := b.Engine().Get(task)
	if err != nil {
		return false, err.Error(), nil
	}
	if !ok {
		return false, "Post the comment failed, no corresponding task", nil
	}

	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}

	// 插入评论
	item := &model.Comment{
		TaskId:    taskId,
		UserId:    strconv.Itoa(ctx.Value("id").(int)),
		Content:   comment,
		Timestamp: time.Now(),
	}

	if _, err := sess.Insert(item); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}

	if err != nil {
		return false, err.Error(), nil
	}

	if _, err = sess.Desc("id").Limit(1, 0).Get(item); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}

	if err := sess.Commit(); err != nil {
		return false, err.Error(), nil
	}
	return true, "", item
}

// 修改评论
func (b *basicCommentService) ChangeComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data *model.Comment) {
	// TODO implement the business logic of ChangeComment
	item := &model.Comment{
		Id:        cId,
		TaskId:    taskId,
		UserId:    strconv.Itoa(ctx.Value("id").(int)),
		Timestamp: time.Now(),
	}
	return status, errinfo, data
}

// 删除评论
func (b *basicCommentService) DeleteComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data string) {
	// TODO implement the business logic of DeleteComment
	return status, errinfo, data
}

// 点赞评论
func (b *basicCommentService) LikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment) {
	// TODO implement the business logic of LikeComment
	return status, errinfo, data
}

// 取消点赞评论
func (b *basicCommentService) CancelLikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment) {
	// TODO implement the business logic of CancelLikeComment
	return status, errinfo, data
}

// NewBasicCommentService returns a naive, stateless implementation of CommentService.
func NewBasicCommentService() CommentService {
	basicCommentSvc := &basicCommentService{
		&db.DBService{},
	}

	err := basicCommentSvc.Bind("conf/conf.lyt.yml")
	if err != nil {
		log.Printf("The CommentService failed to bind with mysql")
	}
	return basicCommentSvc
}

// New returns a CommentService with all of the expected middleware wired in.
func New(middleware []Middleware) CommentService {
	var svc CommentService = NewBasicCommentService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
