package models

import (
	"backend/src/app/models/concerns"
	"backend/src/config"
	"time"
)

type Book struct {
	concerns.Base
	GenreID     uint      `json:"genre_id"`
	Genre       Genre     `json:"genre"`
	PublisherID uint      `json:"publisher_id"`
	Publisher   Publisher `json:"publisher"`
	AuthorID    uint      `json:"author_id"`
	Author      Author    `json:"author"`
	Title       string    `gorm:"size:255" json:"name"`
	ImageUrl    string    `gorm:"size:255" json:"image_url,omitempty"`
	PublishDate time.Time `sql:"type:date" json:"publish_date,omitempty"`
}

func GetAllBooks(books *[]Book) {
	db := config.DbConnect()
	defer db.Close()

	db.Preload("Author").Preload("Genre").Preload("Publisher").Find(&books)
}

func GetBook(book *Book, id string) {
	db := config.DbConnect()
	defer db.Close()

	db.Preload("Author").Preload("Genre").Preload("Publisher").First(&book, id)
}

func InsertBook(book *Book) {
	db := config.DbConnect()
	defer db.Close()

	db.Create(&book)
}

func UpdateBook(book *Book, id string) {
	db := config.DbConnect()
	defer db.Close()

	db.Model(&book).Where("id = ?", id).Updates(
		Book{
			GenreID:     book.GenreID,
			PublisherID: book.PublisherID,
			AuthorID:    book.AuthorID,
			Title:       book.Title,
			ImageUrl:    book.ImageUrl,
			PublishDate: book.PublishDate,
		},
	)
}

func DeleteBook(id string) {
	db := config.DbConnect()
	defer db.Close()

	db.Where("id = ?", id).Delete(&Book{})
}
