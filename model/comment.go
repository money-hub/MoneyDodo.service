package model

import "time"

type Comment struct {
	Id         string    `json:"id" xorm:"<-"`
	TaskId     string    `json:"taskId" xorm:"taskId"`
	UserId     string    `json:"userId" xorm:"userId"`
	Timestamp  time.Time `json:"timestamp" xorm:"timestamp"`
	Content    string    `json:"content" xorm:"content"`
	Stars      int       `json:"stars" xorm:"stars"`
	Stargazers []string  `json:"stargazers" xorm:"stargazers"`
}
