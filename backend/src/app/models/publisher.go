package models

import (
	"backend/src/app/models/concerns"
	"backend/src/config"
)

type Publisher struct {
	concerns.Base
	Name  string `gorm:"size:255" json:"name"`
	Books []Book
}

func GetAllPublishers(publishers *[]Publisher) {
	db := config.DbConnect()
	defer db.Close()

	db.Find(&publishers)
}

func GetPublisher(publisher *Publisher, name string) {
	db := config.DbConnect()
	defer db.Close()

	db.Where("name = ?", name).First(&publisher)
}

func InsertPublisher(publisher *Publisher) {
	db := config.DbConnect()
	defer db.Close()

	db.Create(&publisher)
}
