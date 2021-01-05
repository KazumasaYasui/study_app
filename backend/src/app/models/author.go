package models

import (
	"backend/src/app/models/concerns"
	"backend/src/config"
)

type Author struct {
	concerns.Base
	Name  string `gorm:"size:255" json:"name"`
	Books []Book
}

func GetAllAuthors(authors *[]Author) {
	db := config.DbConnect()
	defer db.Close()

	db.Find(&authors)
}

func GetAuthor(author *Author, name string) {
	db := config.DbConnect()
	defer db.Close()

	db.Where("name = ?", name).First(&author)
}

func InsertAuthor(author *Author) {
	db := config.DbConnect()
	defer db.Close()

	db.Create(&author)
}
