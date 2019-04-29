package model

const (
	DealStateUnderway = "underway"
	DealStateClosure  = "closure"
)

type Deal struct {
	TaskId    string `json:"taskId" xorm:"taskId"`
	Publisher string `json:"publisher" xorm:"publisher"`
	Recipient string `json:"recipient" xorm:"recipient"`
	Since     string `json:"since" xorm:"since"`
	Until     string `json:"until" xorm:"until"`
	Reward    string `json:"reward" xorm:"reward"`
	State     string `json:"state" xorm:"state"`
}
