package models

type KVStorage struct {
	BasicModel
	ProjectId string  `gorm:"type:varchar(300);index" json:"project_id"`
	Account   string  `gorm:"type:varchar(300);index" json:"account"`
	Bucket    string  `gorm:"type:varchar(300);index" json:"bucket"`
	Key       string  `gorm:"type:varchar(600);index" json:"key"`
	Search    string  `gorm:"type:varchar(600);index" json:"search"` // 前後 like 不走索引
	Payload   JSONMap `json:"payload"`
}
