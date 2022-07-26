package models

type KVStorage struct {
	BasicModel
	ProjectId string `gorm:"type:varchar(300);index"`
	Account   string `gorm:"type:varchar(300);index"`
	Bucket    string `gorm:"type:varchar(300);index"`
	Key       string `gorm:"type:varchar(600);index"`
	Search    string `gorm:"type:varchar(600);index"` // 前後 like 不走索引
	Payload   JSONMap
}
