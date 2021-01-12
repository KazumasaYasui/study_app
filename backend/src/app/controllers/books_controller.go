package controllers

import (
	"backend/src/app/controllers/concerns"
	"backend/src/app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var bookReq concerns.BookReq
	if err := json.Unmarshal(reqBody, &bookReq); err != nil {
		log.Fatal(err)
	}

	// Meta情報をFetchし、レコードが存在しない場合はそれぞれInsert
	var author models.Author
	var genre models.Genre
	var publisher models.Publisher
	models.GetAuthor(&author, bookReq.Author)
	models.GetGenre(&genre, bookReq.Genre)
	models.GetPublisher(&publisher, bookReq.Publisher)
	if author.ID == 0 {
		newAuthor := models.Author{Name: bookReq.Author}
		models.InsertAuthor(&newAuthor)
	}
	if genre.ID == 0 {
		newGenre := models.Genre{Name: bookReq.Genre}
		models.InsertGenre(&newGenre)
	}
	if publisher.ID == 0 {
		newPublisher := models.Publisher{Name: bookReq.Publisher}
		models.InsertPublisher(&newPublisher)
	}

	newBook := models.Book{AuthorID: author.ID, GenreID: genre.ID, PublisherID: publisher.ID,
		Title: bookReq.Title, ImageUrl: bookReq.ImageUrl, PublishDate: convertTime(bookReq.PublishDate)}
	models.InsertBook(&newBook)
	respBody, err := json.Marshal(newBook)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func convertTime(str string) time.Time {
	result, _ := time.Parse("2006-01-02", str)
	return result
}
