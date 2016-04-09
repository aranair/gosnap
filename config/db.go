package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func InitDb(dataSourceName string) {
	var err error
	Db, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
}
