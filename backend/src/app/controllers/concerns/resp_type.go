package concerns

import "backend/src/app/models"

type DeleteResp struct {
	Id string `json:"id"`
}

type BookMetaInfoResp struct {
	Authors    []models.Author
	Genres     []models.Genre
	Publishers []models.Publisher
}
