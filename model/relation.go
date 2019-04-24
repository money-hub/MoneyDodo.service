package model

type Relation struct {
	UserId string `json:"userId" xorm:"userId"`
	TaskId string `json:"taskId" xorm:"taskId"`
	Detail string `json:"info" xorm:"Info"` // release or claim
}

func NewEmptyRelation() Relation {
	r := Relation{}
	return r
}

func NewRelation(userId string, taskId string, detail string) Relation {
	r := Relation{
		UserId: userId,
		TaskId: taskId,
		Detail: detail,
	}
	return r
}
