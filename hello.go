package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{names}", HomeHandler).
		Methods("GET")

	http.Handle("/", handlers.LoggingHandler(os.Stdout, router))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
