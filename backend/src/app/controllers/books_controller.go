package controllers

import (
	"backend/src/app/controllers/concerns"
	"backend/src/app/models"
	"encoding/json"
	"log"
	"net/http"
)

func FetchAllBooksMetaInfo(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	var genres []models.Genre
	var publishers []models.Publisher

	models.GetAllAuthors(&authors)
	models.GetAllGenres(&genres)
	models.GetAllPublishers(&publishers)

	booksMetaInfo := concerns.BookMetaInfoResp{Authors: authors, Genres: genres, Publishers: publishers}
	respBody, err := json.Marshal(booksMetaInfo)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}
