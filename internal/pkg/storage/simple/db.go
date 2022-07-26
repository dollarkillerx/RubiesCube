package simple

import (
	"github.com/dollarkillerx/RubiesCube/internal/conf"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"github.com/dollarkillerx/RubiesCube/internal/utils"
	"gorm.io/gorm"
)

type Simple struct {
	db *gorm.DB
}

func NewSimple(conf *conf.PgSQLConfig) (*Simple, error) {
	sql, err := utils.InitPgSQL(conf)
	if err != nil {
		return nil, err
	}

	sql.AutoMigrate(
		&models.ManagerUser{},
		&models.ManagerProject{},
		&models.UserCenter{},
		&models.KVStorage{},
	)

	return &Simple{
		db: sql,
	}, nil
}
