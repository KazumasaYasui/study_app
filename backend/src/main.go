package main

import (
	"backend/src/app/controllers"
	"backend/src/app/controllers/concerns"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello golang from docker!")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/users", concerns.TokenVerifyMiddleWare(controllers.CreateUser)).Methods("POST")
	r.HandleFunc("/users", concerns.TokenVerifyMiddleWare(controllers.FetchAllUsers)).Methods("GET")
	r.HandleFunc("/users/{id}", concerns.TokenVerifyMiddleWare(controllers.FetchUser)).Methods("GET")
	r.HandleFunc("/users/{id}", concerns.TokenVerifyMiddleWare(controllers.UpdateUser)).Methods("PUT")
	r.HandleFunc("/users/{id}", concerns.TokenVerifyMiddleWare(controllers.DeleteUser)).Methods("DELETE")

	r.HandleFunc("/books", concerns.TokenVerifyMiddleWare(controllers.CreateBook)).Methods("POST")
	r.HandleFunc("/books", concerns.TokenVerifyMiddleWare(controllers.FetchAllBooks)).Methods("GET")
	r.HandleFunc("/books/meta_info", concerns.TokenVerifyMiddleWare(controllers.FetchAllBooksMetaInfo)).Methods("GET")
	r.HandleFunc("/books/{id}", concerns.TokenVerifyMiddleWare(controllers.FetchBook)).Methods("GET")
	r.HandleFunc("/books/{id}", concerns.TokenVerifyMiddleWare(controllers.UpdateBook)).Methods("PUT")
	r.HandleFunc("/books/{id}", concerns.TokenVerifyMiddleWare(controllers.DeleteBook)).Methods("DELETE")
	r.HandleFunc("/books/{id}/my_info", concerns.TokenVerifyMiddleWare(controllers.CreateBookMyInfo)).Methods("POST")
	r.HandleFunc("/books/{id}/my_info", concerns.TokenVerifyMiddleWare(controllers.FetchBookMyInfo)).Methods("GET")
	r.HandleFunc("/books/{id}/my_info", concerns.TokenVerifyMiddleWare(controllers.UpdateBookMyInfo)).Methods("PUT")
	r.HandleFunc("/books/{id}/my_info", concerns.TokenVerifyMiddleWare(controllers.DeleteBookMyInfo)).Methods("DELETE")

	log.Println("starting server on 8080 port...")
	http.ListenAndServe(":8080", r)
}
