package model

import "time"

type Charge struct {
	Id        string    `json:"id" xorm:"<-"`
	UserId    string    `json:"userId" xorm:"userId"`
	Amount    float64   `json:"amount" xorm:"amount"`
	Timestamp time.Time `json:"timestamp" xorm:"timestamp"`
}
