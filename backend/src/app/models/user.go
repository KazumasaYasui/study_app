package models

import (
	"backend/src/app/models/concerns"
	"backend/src/config"
)

type User struct {
	concerns.Base
	Name           string         `gorm:"size:255" json:"name"`
	Email          string         `gorm:"size:255" json:"email"`
	PasswordDigest string         `gorm:"size:255" json:"password_digest,omitempty"`
	ImageUrl       string         `gorm:"size:255" json:"image_url,omitempty"`
	UserBookItem   []UserBookItem `json:"-"`
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

func InsertUser(user *User) {
	db := config.DbConnect()
	defer db.Close()

	db.Create(&user)
}

func UpdateUser(user *User, id string) {
	db := config.DbConnect()
	defer db.Close()

	db.Model(&user).Where("id = ?", id).Updates(
		User{
			Email:          user.Email,
			Name:           user.Name,
			PasswordDigest: user.PasswordDigest,
			ImageUrl:       user.ImageUrl,
		},
	)
}

func DeleteUser(id string) {
	db := config.DbConnect()
	defer db.Close()

	db.Where("id = ?", id).Delete(&User{})
}
