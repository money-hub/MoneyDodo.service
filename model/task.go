package model

import "time"

const (
	TaskActionRelease    = "release"
	TaskActionClaim      = "claim"
	TaskActionFinish     = "finish"
	TaskStateNonReleased = "non-released"
	TaskStateReleased    = "released"
	TaskStateClaimed     = "claimed"
	TaskStateFinished    = "finished"
)

type Task struct {
	Id        string     `json:"id" xorm:"<-"`
	Type      string     `json:"type"`
	Publisher string     `json:"publisher"`
	Recipient string     `json:"recipient"`
	Restrain  string     `json:"restrain"`
	Pubdate   *time.Time `json:"pubdate"`
	Cutoff    *time.Time `json:"cutoff"`
	Reward    float64    `json:"reward"`
	State     string     `json:"state"`
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
