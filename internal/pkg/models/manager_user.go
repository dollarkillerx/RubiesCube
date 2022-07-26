package models

type ManagerUser struct {
	BasicModel
	Email    string `gorm:"type:varchar(300);uniqueIndex"`
	Password string `gorm:"type:varchar(600)"`
}
