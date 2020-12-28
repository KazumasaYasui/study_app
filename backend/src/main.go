package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello golang from docker!")
}

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
