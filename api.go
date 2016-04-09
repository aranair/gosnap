package main

import (
	"net/http"

	"github.com/aranair/gosnap/crawler"
	"github.com/aranair/gosnap/models"
	"github.com/gin-gonic/gin"
)

type __api__ struct {
}

func (api __api__) bind(r *gin.RouterGroup) {
	r.GET("/listings", api.listingIndex)
	r.GET("/update_listings", api.updateListings)
}

func (api __api__) listingIndex(c *gin.Context) {
	l := models.ListListings()
	c.JSON(http.StatusOK, gin.H{
		"listings": l,
	})
}

func (api __api__) updateListings(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
