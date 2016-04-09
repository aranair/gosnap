package main

import (
	"log"

	"github.com/aranair/gosnap/config"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"github.com/BurntSushi/toml"
)

func main() {
	var conf config.Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}

	pqStr := "user=" + conf.DB.User + " password='" + conf.DB.Password + "' dbname=gosnap host=localhost sslmode=disable"

	config.InitDb(pqStr)

	router := gin.Default()

	var api = __api__{}

	// Attach api
	api.bind(router.Group(conf.Api.Prefix))

	// For all other requests, see: react.go.
	// react.bind(router)

	// Start listening
	router.Run(":5000")
}
