package model

const (
	ReviewStatePassed   = "passed"
	ReviewStateFailed   = "failed"
	ReviewStateUnderway = "underway"
)

type Review struct {
	Id     string `json:"id" xorm:"<-"`
	TaskId string `json:"taskId" xorm:"taskId"`
	Name   string `json:"name" xorm:"name"`
	State  string `json:"state" xorm:"state"`
}
