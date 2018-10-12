package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GenericMessage default Message
type GenericMessage struct {
	Message string `json: "message"`
}

//Easy message realy structure
type genericResponse struct {
	Data GenericMessage `json: "data"`
}

func main() {
	// Init Router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/h", helloW).Methods("GET")
	log.Fatal(http.ListenAndServe(":23450", router))
}

// Displays hello world
func helloW(w http.ResponseWriter, r *http.Request) {
	response := genericResponse{
		Data: GenericMessage{
			Message: "Hello World",
		},
	}

	json.NewEncoder(w).Encode(response)
}
