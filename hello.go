package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{names}", HomeHandler)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	names := vars["names"]

	fmt.Fprintf(response, "Hi %s! Nice to see you!", names)
}
