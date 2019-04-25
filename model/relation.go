package model

type Relation struct {
	UserId string `json:"userId" xorm:"userId"`
	TaskId string `json:"taskId" xorm:"taskId"`
	Status string `json:"status" xorm:"status"` // release or claim
}
