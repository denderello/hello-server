package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	names := vars["names"]

	fmt.Fprintf(response, "Hi %s! Nice to see you!", names)
}
