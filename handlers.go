package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	names := splitNamesParameter(string(vars["names"]))

	fmt.Fprintf(response, "Hi %s! Nice to see you!", concatNames(names))
}

func splitNamesParameter(concatedNames string) []string {
	return strings.Split(concatedNames, " ")
}

func concatNames(names []string) string {
	return strings.Join(names, ", ")
}
