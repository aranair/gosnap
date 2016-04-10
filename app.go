package main

import (
	"log"

	"github.com/aranair/gosnap/config"
	"github.com/aranair/gosnap/crawler"
	"github.com/aranair/gosnap/models"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"github.com/BurntSushi/toml"
	"github.com/bamzi/jobrunner"
)

func main() {
	var conf config.Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}

	jobrunner.Start()
	jobrunner.Schedule("@every 30s", UpdateListings{})

	pqStr := "user=" + conf.DB.User + " password='" + conf.DB.Password + "' dbname=gosnap host=localhost sslmode=disable"

	config.InitDb(pqStr)
	router := gin.Default()

	var api = __api__{}

	// Attach api
	api.bind(router.Group(conf.Api.Prefix))

	// router.GET("/jobrunner/json", JobJson)
	// routes.LoadHTMLGlob("../github.com/bamzi/jobrunner/views/Status.html")
	// routes.GET("/jobrunner/html", JobHtml)
	router.Run(":5000")
}

type UpdateListings struct{}

func (e UpdateListings) Run() {
	urls := map[int][]string{
		1: []string{"http://www.clubsnap.com/forums/forumdisplay.php?f=104"},
		2: []string{"http://www.clubsnap.com/forums/forumdisplay.php?f=102"},
		3: []string{"http://www.clubsnap.com/forums/forumdisplay.php?f=111"},
	}

	for cid, urla := range urls {
		l := crawler.Start(urla)
		for title, url := range l {
			models.CreateListing(cid, title, url)
		}
	}
}
