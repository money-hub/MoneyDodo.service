package model

import "time"

type Recharge struct {
	Id        string    `json:"id" xorm:"<-"`
	UserId    string    `json:"userId" xorm:"userId"`
	Amount    string    `json:"rechargeAmount" xorm:"rechargeAmount"`
	Timestamp time.Time `json:"timestamp" xorm:"timestamp"`
}
