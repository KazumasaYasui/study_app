package models

import (
	"backend/src/app/models/concerns"
	"backend/src/config"
)

type User struct {
	concerns.Base
	Name           string `gorm:"size:255" json:"name"`
	Email          string `gorm:"size:255" json:"email"`
	PasswordDigest string `json:"password_digest,omitempty"`
	ImageUrl       string `json:"image_url,omitempty"`
}

func GetAllUsers(users *[]User) {
	db := config.DbConnect()
	defer db.Close()

	db.Find(&users)
}

func GetUser(user *User, id string) {
	db := config.DbConnect()
	defer db.Close()

	db.First(&user, id)
}
