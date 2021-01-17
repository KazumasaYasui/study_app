package controllers

import (
	"backend/src/app/controllers/concerns"
	"backend/src/app/models"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var signupReq concerns.SignupReq
	if err := json.Unmarshal(reqBody, &signupReq); err != nil {
		log.Fatal(err)
	}

	if signupReq.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(concerns.ErrorResp{Message: "Emailは必須です"})
		w.Write(respBody)
		return
	}
	if signupReq.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(concerns.ErrorResp{Message: "Passwordは必須です"})
		w.Write(respBody)
		return
	}
	if signupReq.Password != signupReq.PasswordConfirmation {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(concerns.ErrorResp{Message: "Passwordが一致しません"})
		w.Write(respBody)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), 10)
	if err != nil {
		log.Fatal(err)
	}

	tmpName, err := makeRandStr(12)
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{Email: signupReq.Email, PasswordDigest: string(hash), Name: tmpName}
	models.InsertUser(&user)
	respBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func makeRandStr(digit int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
