package controllers

import (
	"backend/src/app/controllers/concerns"
	"backend/src/app/models"
	"backend/src/config"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var loginReq concerns.LoginReq
	if err1 := json.Unmarshal(reqBody, &loginReq); err1 != nil {
		log.Fatal(err1)
	}

	if loginReq.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(concerns.ErrorResp{Message: "Emailは必須です"})
		w.Write(respBody)
		return
	}
	if loginReq.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(concerns.ErrorResp{Message: "Passwordは必須です"})
		w.Write(respBody)
		return
	}

	var user models.User
	models.GetUserWithEmail(&user, loginReq.Email)
	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		respBody, _ := json.Marshal(concerns.ErrorResp{Message: "登録されていません"})
		w.Write(respBody)
		return
	}

	password := loginReq.Password
	hashedPassword := user.PasswordDigest
	err2 := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err2 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		respBody, _ := json.Marshal(concerns.ErrorResp{Message: "無効なパスワードです"})
		w.Write(respBody)
		return
	}

	token, err3 := createToken(user)
	if err3 != nil {
		log.Fatal(err3)
	}

	err4 := setSession(strconv.Itoa(int(user.ID)), token)
	if err4 != nil {
		log.Fatal(err4)
	}

	respBody, _ := json.Marshal(concerns.TokenResp{Token: token})

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func createToken(user models.User) (string, error) {
	secret := "tmpSecretKey"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "KazumasaYasui",
		"iat":   time.Now(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

func setSession(userId, token string) error {
	ctx := context.Background()
	client := config.SessionConnect()
	err := client.Set(ctx, userId, token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
