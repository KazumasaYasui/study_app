package main

import (
	"backend/src/app/controllers"
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
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.FetchAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.FetchUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books", controllers.FetchAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.FetchBook).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/meta_info", controllers.FetchAllBooksMetaInfo).Methods("GET")
	r.HandleFunc("/books/{id}/my_info", controllers.CreateBookMyInfo).Methods("POST")
	r.HandleFunc("/books/{id}/my_info", controllers.FetchBookMyInfo).Methods("GET")
	r.HandleFunc("/books/{id}/my_info", controllers.UpdateBookMyInfo).Methods("PUT")
	r.HandleFunc("/books/{id}/my_info", controllers.DeleteBookMyInfo).Methods("DELETE")

	log.Println("starting server on 8080 port...")
	http.ListenAndServe(":8080", r)
}
