package main

import (
	"database/sql"
	"net/http"

	"github.com/aranair/gosnap/listings"
	"github.com/gin-gonic/gin"
)

type __api__ struct {
	DB *sql.DB
}

func (api __api__) bind(r *gin.RouterGroup) {
	r.GET("/listings", api.listingIndex)
	r.GET("/update_listings", api.updateListings)
}

func (api __api__) listingIndex(c *gin.Context) {
	var listings = listings.AllListings(api.DB)
	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
	})
}

func (api __api__) updateListings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
