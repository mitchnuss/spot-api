package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
)

type People struct {
	CreatedAt      int64  `dynamodb:"CreatedAt" json:"CreatedAt"`
	Email          string `dynamodb:"Email" json:"Email"`
	FirstDecision  string `dynamodb:"FirstDecision" json:"FirstDecision"`
	FirstName      string `dynamodb:"FirstName" json:FirstName"`
	FullName       string `dynamodb:"Fullname" json:FullName"`
	LastName       string `dynamodb:"LastName" json:LastName"`
	LastUpdated    int64  `dynamodb:"LastUpdated" json:LastUpdated"`
	MembershipType string `dynamodb:"MembershipType" json:MembershipType"`
	NewCreation    string `dynamodb:"NewCreation" json:NewCreation"`
	PhoneNumber    int    `dynamodb:"PhoneNumber" json:PhoneNumber"`
	Rededication   string `dynamodb:"Rededication" json:Rededication"`
	UUID           string `dynamodb:"UUID" json:UUID"`
}

var people []People

func main() {
	// Init Router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/people/{id}", getSingle).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Will pull one person's info
func getSingle(w http.ResponseWriter, r *http.Request) {

	//starts session with us-east-1 dynamodb and loads credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	//creates Dynamo Client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("SPOT"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {
				N: aws.String("id"),
			},
		},
	})

	// loops through UUID's and encodes to json
	for _, item := range people {
		if item.UUID == People["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&People{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	item := People{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	} //Sprintf formats according to a format specifier and returns the resulting string.

	if item.UUID == "" {
		fmt.Println("Could not find the UUID specified")
	}
}
