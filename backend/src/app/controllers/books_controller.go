package controllers

import (
	"backend/src/app/controllers/concerns"
	"backend/src/app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FetchAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	models.GetAllBooks(&books)
	respBody, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func FetchBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var book models.Book

	models.GetBook(&book, id)
	respBody, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

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
