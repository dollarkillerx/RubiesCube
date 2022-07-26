package main

import (
	"fmt"
	"log"

	"github.com/dollarkillerx/RubiesCube/internal/conf"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/storage/simple"
	"github.com/rs/xid"
	"gorm.io/datatypes"
)

func main() {
	simple, err := simple.NewSimple(&conf.CONF.PgSQLConfig)
	if err != nil {
		log.Fatalln(err)
	}

	//Ins(simple)

	planets := []models.KVStorage{}
	// 存在Key
	simple.DB().Find(&planets, datatypes.JSONQuery("payload").Equals("gaia", "alt_name").Equals("ok", "jx"))
	fmt.Printf("%+v\n", planets)
	fmt.Printf("%+v\n", len(planets))

	//planets = []models.KVStorage{}
	//// 存在Key
	//simple.DB().Find(&planets, datatypes.JSONQuery("payload").HasKey("is_red"))
	//fmt.Printf("%+v", planets)
}

func Ins(simple *simple.Simple) {
	earthSpecs := map[string]interface{}{
		"alt_name":         "gaia",
		"jx":               "ok",
		"is_red":           false,
		"main_composition": []string{"nitrogen", "oxygen"},
		"others": map[string]interface{}{
			"is_nice":             true,
			"orbital_period_days": 365,
		},
	}

	err := simple.DB().
		Model(&models.KVStorage{}).
		Create(&models.KVStorage{
			BasicModel: models.BasicModel{
				ID: xid.New().String(),
			},
			ProjectId: "xssx",
			Account:   "asdasd",
			Key:       "Earth",
			Payload:   earthSpecs,
		}).Error
	if err != nil {
		log.Fatalln(err)
	}

	marsSpecs := map[string]interface{}{
		"alt_name": "gaia2x",
		"main_composition": []string{
			"carbon dioxide",
			"argon",
			"nitrogen",
			"oxygen",
		},
		"is_red": true,
		"others": map[string]interface{}{
			"is_nice":             "hum...",
			"orbital_period_days": 686,
		},
	}

	err = simple.DB().
		Model(&models.KVStorage{}).
		Create(&models.KVStorage{
			BasicModel: models.BasicModel{
				ID: xid.New().String(),
			},
			ProjectId: "xssx",
			Account:   "asdasd",
			Key:       "Mars",
			Payload:   marsSpecs,
		}).Error
	if err != nil {
		log.Fatalln(err)
	}

	jupiterSpecs := map[string]interface{}{
		"alt_name":         "gaia3x",
		"main_composition": []string{"hydrogen", "helium"},
	}
	err = simple.DB().
		Model(&models.KVStorage{}).
		Create(&models.KVStorage{
			BasicModel: models.BasicModel{
				ID: xid.New().String(),
			},
			ProjectId: "xssx",
			Account:   "asdasd",
			Key:       "Jupiter",
			Payload:   jupiterSpecs,
		}).Error
	if err != nil {
		log.Fatalln(err)
	}
}
