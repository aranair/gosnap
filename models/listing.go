package models

import (
	"time"

	"github.com/aranair/gosnap/config"
	"github.com/jinzhu/gorm"
)

type Listing struct {
	gorm.Model
	Title      string `json:"title"`
	Url        string `json:"url"`
	CategoryId int    `json:"cateogry_id"`
}

func ListListings() (list []Listing) {
	config.Db.Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).Find(&list)
	return
}

func CreateListing(cid int, title string, url string) (listingId int) {
	l := Listing{Title: title, Url: url, CategoryId: cid}
	config.Db.Create(&l)
	return
}
