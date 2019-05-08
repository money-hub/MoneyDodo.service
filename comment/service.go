package comment

// CommentService provide comments related operations

import (
	"context"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

type CommentService interface {

	// 获取某个task的评论
	GetComment(ctx context.Context, taskId string) (status bool, errinfo string, data []model.Comment)

	// 发表评论
	PostComment(ctx context.Context, taskId string, comment string) (status bool, errinfo string, data string)

	// 更改某条评论
	ChangeComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data string)

	// 删除某条评论
	DeleteComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data string)

	// 点赞某条评论
	LikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data string)

	// 取消点赞某条评论
	CancelLikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data string)
}

type basicCptService struct {
	*db.DBService
}
