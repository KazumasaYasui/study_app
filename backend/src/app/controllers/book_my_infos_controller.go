package controllers

import (
	"backend/src/app/controllers/concerns"
	"backend/src/app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func CreateBookMyInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId := "1" // TODO: 認証機能作成後、トークンからuserIdを取得する予定
	uintId, _ := strconv.ParseUint(id, 10, 64)
	uintUserId, _ := strconv.ParseUint(userId, 10, 64)
	reqBody, _ := ioutil.ReadAll(r.Body)

	var bookMyInfoReq concerns.BookMyInfoReq
	if err := json.Unmarshal(reqBody, &bookMyInfoReq); err != nil {
		log.Fatal(err)
	}

	item := models.UserBookItem{BookID: uint(uintId), UserID: uint(uintUserId), Status: bookMyInfoReq.Status,
		ProgressRate: bookMyInfoReq.ProgressRate, Urgency: bookMyInfoReq.Urgency, Priority: bookMyInfoReq.Priority}
	models.InsertBookMyInfo(&item)
	respBody, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func UpdateBookMyInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId := "1" // TODO: 認証機能作成後、トークンからuserIdを取得する予定
	uintId, _ := strconv.ParseUint(id, 10, 64)
	uintUserId, _ := strconv.ParseUint(userId, 10, 64)
	reqBody, _ := ioutil.ReadAll(r.Body)

	var bookMyInfoReq concerns.BookMyInfoReq
	if err := json.Unmarshal(reqBody, &bookMyInfoReq); err != nil {
		log.Fatal(err)
	}

	item := models.UserBookItem{BookID: uint(uintId), UserID: uint(uintUserId), Status: bookMyInfoReq.Status,
		ProgressRate: bookMyInfoReq.ProgressRate, Urgency: bookMyInfoReq.Urgency, Priority: bookMyInfoReq.Priority}
	models.UpdateBookMyInfo(&item, id, userId)
	convertUintId, _ := strconv.ParseUint(bookMyInfoReq.Id, 10, 64)
	item.Base.ID = uint(convertUintId)
	respBody, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func DeleteBookMyInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId := "1" // TODO: 認証機能作成後、トークンからuserIdを取得する予定

	models.DeleteBookMyInfo(id, userId)
	respBody, err := json.Marshal(concerns.DeleteResp{BookId: id, UserId: userId})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}
