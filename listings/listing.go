package listings

import (
	"database/sql"
	"log"
)

type Listing struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func AllListings(DB *sql.DB) []Listing {
	list := []Listing{}
	rows, err := DB.Query("SELECT title, url FROM listings limit 50")

	for rows.Next() {
		var l Listing
		err = rows.Scan(&l.Title, &l.Url)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		list = append(list, l)
	}

	return list
}

func InsertListing(DB *sql.DB, title string, url string) (listingId int) {
	DB.QueryRow(`INSERT INTO listings(title, url)
    VALUES($1, $2) RETURNING id`, title, url).Scan(&listingId)

	return
}
