package models

import (
	"backend/src/app/models/concerns"
	"backend/src/config"
)

type Genre struct {
	concerns.Base
	Name  string `gorm:"size:255" json:"name"`
	Books []Book
}

func GetAllGenres(genres *[]Genre) {
	db := config.DbConnect()
	defer db.Close()

	db.Find(&genres)
}

func GetGenre(genre *Genre, name string) {
	db := config.DbConnect()
	defer db.Close()

	db.Where("name = ?", name).First(&genre)
}

func InsertGenre(genre *Genre) {
	db := config.DbConnect()
	defer db.Close()

	db.Create(&genre)
}
