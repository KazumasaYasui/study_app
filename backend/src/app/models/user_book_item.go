package models

import (
	"backend/src/app/models/concerns"
	"backend/src/config"
)

type UserBookItem struct {
	concerns.Base
	UserID       uint    `json:"user_id"`
	User         User    `json:"-"`
	BookID       uint    `json:"book_id"`
	Book         Book    `json:"-"`
	Status       string  `json:"status"`
	ProgressRate float64 `json:"progress_rate"`
	Urgency      int     `json:"urgency"`
	Priority     int     `json:"priority"`
}

func GetBookMyInfo(item *UserBookItem, bookId, userId string) {
	db := config.DbConnect()
	defer db.Close()

	db.Preload("User").Preload("Book").Where(
		"book_id = ? AND user_id = ?", bookId, userId).First(&item)
}
