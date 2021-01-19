package concerns

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("不正な署名方式です")
				}
				return []byte("tmpSecretKey"), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				respBody, _ := json.Marshal(ErrorResp{Message: "tokenデコード時にエラーが発生しました"})
				w.Write(respBody)
				return
			}

			if token.Valid {
				ctx := r.Context()
				authEmail := token.Claims.(jwt.MapClaims)["email"]
				ctx = context.WithValue(ctx, "authEmail", authEmail)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				respBody, _ := json.Marshal(ErrorResp{Message: "tokenが無効です"})
				w.Write(respBody)
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			respBody, _ := json.Marshal(ErrorResp{Message: "tokenが不正です"})
			w.Write(respBody)
			return
		}
	})
}
