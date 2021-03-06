package model

type User struct {
	Id                  string  `json:"id"`
	SId                 string  `json:"sId" xorm:"sId"`
	Name                string  `json:"name"`
	Introduction        string  `json:"introduction"`
	Balance             float64 `json:"balance"`
	Icon                string  `json:"icon"`
	Phone               string  `json:"phone"`
	CreditScore         int     `json:"creditScore" xorm:"creditScore"`
	Email               string  `json:"email"`
	CertifiedPic        string  `json:"certifiedPic" xorm:"certifiedPic"`
	CertificationStatus int     `json:"certificationStatus" xorm:"certificationStatus"`
}
