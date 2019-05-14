package model

import "time"

const (
	DealStateUnderway = "underway"
	DealStateClosure  = "closure"
)

type Deal struct {
	Id        string    `json:"id" xorm:"<-"`
	TaskId    string    `json:"taskId" xorm:"taskId"`
	Publisher string    `json:"publisher" xorm:"publisher"`
	Recipient string    `json:"recipient" xorm:"recipient"`
	Since     time.Time `json:"since" xorm:"since"`
	Until     time.Time `json:"until" xorm:"until"`
	Reward    float64   `json:"reward" xorm:"reward"`
	State     string    `json:"state" xorm:"state"`
}
