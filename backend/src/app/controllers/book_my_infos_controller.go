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
	ctx := r.Context()
	authEmail := ctx.Value("authEmail").(string)
	var user models.User
	models.GetUserWithEmail(&user, authEmail)

	vars := mux.Vars(r)
	id := vars["id"]
	userId := strconv.FormatUint(uint64(user.ID), 10)

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
	ctx := r.Context()
	authEmail := ctx.Value("authEmail").(string)
	var user models.User
	models.GetUserWithEmail(&user, authEmail)

	vars := mux.Vars(r)
	id := vars["id"]
	uintId, _ := strconv.ParseUint(id, 10, 64)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var bookMyInfoReq concerns.BookMyInfoReq
	if err := json.Unmarshal(reqBody, &bookMyInfoReq); err != nil {
		log.Fatal(err)
	}

	item := models.UserBookItem{BookID: uint(uintId), UserID: user.ID, Status: bookMyInfoReq.Status,
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
	ctx := r.Context()
	authEmail := ctx.Value("authEmail").(string)
	var user models.User
	models.GetUserWithEmail(&user, authEmail)

	vars := mux.Vars(r)
	id := vars["id"]
	uintId, _ := strconv.ParseUint(id, 10, 64)
	userId := strconv.FormatUint(uint64(user.ID), 10)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var bookMyInfoReq concerns.BookMyInfoReq
	if err := json.Unmarshal(reqBody, &bookMyInfoReq); err != nil {
		log.Fatal(err)
	}

	item := models.UserBookItem{BookID: uint(uintId), UserID: user.ID, Status: bookMyInfoReq.Status,
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
	ctx := r.Context()
	authEmail := ctx.Value("authEmail").(string)
	var user models.User
	models.GetUserWithEmail(&user, authEmail)

	vars := mux.Vars(r)
	id := vars["id"]
	userId := strconv.FormatUint(uint64(user.ID), 10)

	models.DeleteBookMyInfo(id, userId)
	respBody, err := json.Marshal(concerns.DeleteResp{BookId: id, UserId: userId})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}
