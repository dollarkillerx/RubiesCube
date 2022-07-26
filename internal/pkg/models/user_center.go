package models

type UserCenter struct {
	BasicModel
	ProjectId string `gorm:"type:varchar(300);index"`
	Account   string `gorm:"type:varchar(300);index"`
	Password  string `gorm:"type:varchar(600)"`
	Payload   JSONMap
}
