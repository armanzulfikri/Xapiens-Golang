package main

import (
	"Scraping/config"
	"Scraping/controllers"
	"Scraping/models"
)

func main() {
	//postgre Connect
	pgDB := config.Connect()
	strDB := controllers.StrDB{DB: pgDB}

	// migration
	models.Migrations(pgDB)

	// r := gin.Default()
	//route Comments
	strDB.PostDataFromCommentsAPI()
	strDB.GetComments()

	//route Albums
	strDB.PostDataFromAlbumAPI()
	strDB.GetAlbums()

	//route Photos
	strDB.PostDataFromPhotosAPI()
	strDB.GetPhotos()

	// r.Run()
}
