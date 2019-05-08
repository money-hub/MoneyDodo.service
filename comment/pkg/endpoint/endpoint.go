package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/money-hub/MoneyDodo.service/comment/pkg/service"
	model "github.com/money-hub/MoneyDodo.service/model"
)

// GetCommentRequest collects the request parameters for the GetComment method.
type GetCommentRequest struct {
	TaskId string `json:"task_id"`
}

// GetCommentResponse collects the response parameters for the GetComment method.
type GetCommentResponse struct {
	Status  bool            `json:"status"`
	Errinfo string          `json:"errinfo"`
	Data    []model.Comment `json:"data"`
}

// MakeGetCommentEndpoint returns an endpoint that invokes GetComment on the service.
func MakeGetCommentEndpoint(s service.CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCommentRequest)
		status, errinfo, data := s.GetComment(ctx, req.TaskId)
		return GetCommentResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostCommentRequest collects the request parameters for the PostComment method.
type PostCommentRequest struct {
	TaskId  string `json:"task_id"`
	Comment string `json:"comment"`
}

// PostCommentResponse collects the response parameters for the PostComment method.
type PostCommentResponse struct {
	Status  bool           `json:"status"`
	Errinfo string         `json:"errinfo"`
	Data    *model.Comment `json:"data"`
}

// MakePostCommentEndpoint returns an endpoint that invokes PostComment on the service.
func MakePostCommentEndpoint(s service.CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostCommentRequest)
		status, errinfo, data := s.PostComment(ctx, req.TaskId, req.Comment)
		return PostCommentResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// ChangeCommentRequest collects the request parameters for the ChangeComment method.
type ChangeCommentRequest struct {
	TaskId  string `json:"task_id"`
	CId     string `json:"cid"`
	Comment string `json:"comment"`
}

// ChangeCommentResponse collects the response parameters for the ChangeComment method.
type ChangeCommentResponse struct {
	Status  bool           `json:"status"`
	Errinfo string         `json:"errinfo"`
	Data    *model.Comment `json:"data"`
}

// MakeChangeCommentEndpoint returns an endpoint that invokes ChangeComment on the service.
func MakeChangeCommentEndpoint(s service.CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeCommentRequest)
		status, errinfo, data := s.ChangeComment(ctx, req.TaskId, req.CId, req.Comment)
		return ChangeCommentResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// DeleteCommentRequest collects the request parameters for the DeleteComment method.
type DeleteCommentRequest struct {
	TaskId  string `json:"task_id"`
	CId     string `json:"cid"`
	Comment string `json:"comment"`
}

// DeleteCommentResponse collects the response parameters for the DeleteComment method.
type DeleteCommentResponse struct {
	Status  bool   `json:"status"`
	Errinfo string `json:"errinfo"`
	Data    string `json:"data"`
}

// MakeDeleteCommentEndpoint returns an endpoint that invokes DeleteComment on the service.
func MakeDeleteCommentEndpoint(s service.CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCommentRequest)
		status, errinfo, data := s.DeleteComment(ctx, req.TaskId, req.CId, req.Comment)
		return DeleteCommentResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// LikeCommentRequest collects the request parameters for the LikeComment method.
type LikeCommentRequest struct {
	TaskId string `json:"task_id"`
	CId    string `json:"cid"`
}

// LikeCommentResponse collects the response parameters for the LikeComment method.
type LikeCommentResponse struct {
	Status  bool           `json:"status"`
	Errinfo string         `json:"errinfo"`
	Data    *model.Comment `json:"data"`
}

// MakeLikeCommentEndpoint returns an endpoint that invokes LikeComment on the service.
func MakeLikeCommentEndpoint(s service.CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LikeCommentRequest)
		status, errinfo, data := s.LikeComment(ctx, req.TaskId, req.CId)
		return LikeCommentResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// CancelLikeCommentRequest collects the request parameters for the CancelLikeComment method.
type CancelLikeCommentRequest struct {
	TaskId string `json:"task_id"`
	CId    string `json:"cid"`
}

// CancelLikeCommentResponse collects the response parameters for the CancelLikeComment method.
type CancelLikeCommentResponse struct {
	Status  bool           `json:"status"`
	Errinfo string         `json:"errinfo"`
	Data    *model.Comment `json:"data"`
}

// MakeCancelLikeCommentEndpoint returns an endpoint that invokes CancelLikeComment on the service.
func MakeCancelLikeCommentEndpoint(s service.CommentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CancelLikeCommentRequest)
		status, errinfo, data := s.CancelLikeComment(ctx, req.TaskId, req.CId)
		return CancelLikeCommentResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetComment implements Service. Primarily useful in a client.
func (e Endpoints) GetComment(ctx context.Context, taskId string) (status bool, errinfo string, data []model.Comment) {
	request := GetCommentRequest{TaskId: taskId}
	response, err := e.GetCommentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetCommentResponse).Status, response.(GetCommentResponse).Errinfo, response.(GetCommentResponse).Data
}

// PostComment implements Service. Primarily useful in a client.
func (e Endpoints) PostComment(ctx context.Context, taskId string, comment string) (status bool, errinfo string, data *model.Comment) {
	request := PostCommentRequest{
		Comment: comment,
		TaskId:  taskId,
	}
	response, err := e.PostCommentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostCommentResponse).Status, response.(PostCommentResponse).Errinfo, response.(PostCommentResponse).Data
}

// ChangeComment implements Service. Primarily useful in a client.
func (e Endpoints) ChangeComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data *model.Comment) {
	request := ChangeCommentRequest{
		CId:     cId,
		Comment: comment,
		TaskId:  taskId,
	}
	response, err := e.ChangeCommentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ChangeCommentResponse).Status, response.(ChangeCommentResponse).Errinfo, response.(ChangeCommentResponse).Data
}

// DeleteComment implements Service. Primarily useful in a client.
func (e Endpoints) DeleteComment(ctx context.Context, taskId string, cId string, comment string) (status bool, errinfo string, data string) {
	request := DeleteCommentRequest{
		CId:     cId,
		Comment: comment,
		TaskId:  taskId,
	}
	response, err := e.DeleteCommentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteCommentResponse).Status, response.(DeleteCommentResponse).Errinfo, response.(DeleteCommentResponse).Data
}

// LikeComment implements Service. Primarily useful in a client.
func (e Endpoints) LikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment) {
	request := LikeCommentRequest{
		CId:    cId,
		TaskId: taskId,
	}
	response, err := e.LikeCommentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LikeCommentResponse).Status, response.(LikeCommentResponse).Errinfo, response.(LikeCommentResponse).Data
}

// CancelLikeComment implements Service. Primarily useful in a client.
func (e Endpoints) CancelLikeComment(ctx context.Context, taskId string, cId string) (status bool, errinfo string, data *model.Comment) {
	request := CancelLikeCommentRequest{
		CId:    cId,
		TaskId: taskId,
	}
	response, err := e.CancelLikeCommentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CancelLikeCommentResponse).Status, response.(CancelLikeCommentResponse).Errinfo, response.(CancelLikeCommentResponse).Data
}
