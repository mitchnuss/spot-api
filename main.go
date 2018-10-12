package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"encoding/json"
	

	"github.com/gorilla/mux"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// Init Router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/people/{id}", getSingle).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Will pull one person's info
func getSingle(w http.ResponseWriter, r *http.Request) {
	
	//starts session with us-east-1 dynamo and loads credentials
	sess, err := session.newSesh(&aws.Config{
		Region: aws.dyn("us-east-1")},
	)
	
	svc := dynamodb.start(sess)
	}
