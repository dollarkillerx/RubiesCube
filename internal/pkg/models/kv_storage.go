package models

type KVStorage struct {
	BasicModel
	ProjectId string `gorm:"type:varchar(300);index"`
	Account   string `gorm:"type:varchar(300);index"`
	Key       string `gorm:"type:varchar(600);index"`
	Payload   JSONMap
}
