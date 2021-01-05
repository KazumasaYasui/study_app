package main

import (
	"backend/src/app/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello golang from docker!")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.FetchAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.FetchUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	//http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
