package test

import (
	"github.com/dollarkillerx/RubiesCube/internal/conf"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"github.com/dollarkillerx/RubiesCube/internal/storage/simple"
	"github.com/rs/xid"
	"log"
	"testing"
)

func TestPx(t *testing.T) {
	simple, err := simple.NewSimple(&conf.CONF.PgSQLConfig)
	if err != nil {
		log.Fatalln(err)
	}

	simple.DB().Model(&models.ManagerProject{}).
		Create(&models.ManagerProject{
			BasicModel:  models.BasicModel{ID: xid.New().String()},
			UserID:      "10086",
			ProjectName: "inventory",
			Payload: models.ProjectPayload{
				Buckets: []string{
					"commodity",                // 商品
					"inbound_outbound_records", // 出入庫記錄
				},
			},
		})
}
