package storage

import (
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"gorm.io/gorm"
)

type Interface interface {
	DB() *gorm.DB
	GetUserCenter(projectID string, account string) (*models.UserCenter, error)
}
