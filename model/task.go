package model

import "time"

const (
	TaskStateNonReleased  = "non-released"
	TaskStateReleased     = "released"
	TaskStateClosed       = "closed"
	TaskKindQuestionnaire = "questionnaire"
)

type Task struct {
	Id        string `json:"id" xorm:"<-"`
	Kind      string `json:"kind"`
	Publisher string `json:"publisher"`
	// Recipient       string    `json:"recipient"`
	Restrain string    `json:"restrain"`
	Pubdate  time.Time `json:"pubdate"`
	Cutoff   time.Time `json:"cutoff"`
	// Enddate         time.Time `json:"enddate"`
	Reward float64 `json:"reward"`
	// RecipientFinish bool      `json:"recipientFinish" xorm:"recipientFinish"`
	// ConfirmFinish   bool      `json:"confirmFinish" xorm:"confirmFinish"`
	State string `json:"state"`
}

type query struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type singleChoice struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   string   `json:"answer"`
}

type mutipleChoice struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answers  []string `json:"answers"`
}

type Questionnaire struct {
	TaskId        string          `json:"taskId" xorm:"taskId"`
	Query         []query         `json:"query" xorm:"query"`
	SingleChoice  []singleChoice  `json:"singleChoice" xorm:"singleChoice"`
	MutipleChoice []mutipleChoice `json:"mutipleChoice" xorm:"mutipleChoice"`
}

type Qtnr struct {
	Task
	Qtnr *Questionnaire `json:"qtnr"`
}
