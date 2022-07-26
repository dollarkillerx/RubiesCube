package storage

import (
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"gorm.io/gorm"
)

type Interface interface {
	DB() *gorm.DB
	GetUserCenter(projectID string, account string) (*models.UserCenter, error)
	StorageKV(kvStorage models.KVStorage) error
	KvUpdate(project string, account string, bucket string, key string, payload models.JSONMap) (err error)
	KvDelete(project string, account string, bucket string, key string) (err error)
}
