package controllers

import (
	"backend/src/app/controllers/concerns"
	"backend/src/app/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func FetchAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	models.GetAllUsers(&users)
	respBody, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User

	models.GetUser(&user, id)
	respBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		log.Fatal(err)
	}

	models.InsertUser(&user)
	respBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		log.Fatal(err)
	}

	models.UpdateUser(&user, id)
	convertUintId, _ := strconv.ParseUint(id, 10, 64)
	user.Base.ID = uint(convertUintId)
	respBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	models.DeleteUser(id)
	respBody, err := json.Marshal(concerns.DeleteResp{Id: id})
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}
