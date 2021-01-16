package concerns

import "backend/src/app/models"

type DeleteResp struct {
	Id     string `json:"id,omitempty"`
	BookId string `json:"book_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

type BookMetaInfoResp struct {
	Authors    []models.Author
	Genres     []models.Genre
	Publishers []models.Publisher
}
