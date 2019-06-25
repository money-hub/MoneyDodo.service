package model

const (
	ReviewStatePassed   = "passed"
	ReviewStateFailed   = "failed"
	ReviewStateUnderway = "underway"
)

type Review struct {
	Id     string `json:"id" xorm:"<-"`
	TaskId string `json:"taskId" xorm:"taskId"`
	UserId string `json:"userId" xorm:"userId"`
	State  string `json:"state" xorm:"state"`
}
