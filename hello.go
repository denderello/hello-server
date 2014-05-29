package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{names}", HomeHandler).
		Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
