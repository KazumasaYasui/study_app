package controllers

import (
	"backend/src/app/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FetchBookMyInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId := "1" // TODO: 認証機能作成後、トークンからuserIdを取得する予定
	var item models.UserBookItem

	models.GetBookMyInfo(&item, id, userId)
	respBody, err := json.Marshal(item)
	log.Println(item)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}
