package model

import "time"

type Task struct {
	Id        string     `json:"id" xorm:"<-"`
	Type      string     `json:"type"`
	From      string     `json:"from"`
	Recipient string     `json:"recipient"`
	Limit     string     `json:"limit"`
	Release   *time.Time `json:"release"`
	Cutoff    *time.Time `json:"cutoff"`
	Reward    float64    `json:"reward"`
	Status    string     `json:"status"`
}

type query struct {
	Question string `json:"question"`
	Answer   string `json:"answaer"`
}

type singleChoice struct {
	Question string `json:"question"`
	Choice1  string `json:"choice1"`
	Choice2  string `json:"choice2"`
	Choice3  string `json:"choice3"`
	Choice4  string `json:"choice4"`
	Answer   string `json:"answer"`
}

type Questionnaire struct {
	Task
	Query        []query
	SingleChoice []singleChoice
}
