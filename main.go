package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":23450", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %q", html.EscapeString(r.URL.Path))
}
