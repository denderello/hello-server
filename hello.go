package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var version string = "dev"

func main() {
	fmt.Println("Hello Server version", version)

	router := mux.NewRouter()
	router.HandleFunc("/{names}", HomeHandler).
		Methods("GET")

	http.Handle("/", handlers.LoggingHandler(os.Stdout, router))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
