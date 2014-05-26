package main

import "fmt"
import "log"
import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi %s! Nice to see you!", r.URL.Path[1:])
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
