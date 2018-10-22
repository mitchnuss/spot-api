package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//People information in SPOT
type People struct {
	CreatedAt      int64  `dynamodb:"CreatedAt" json:"CreatedAt"`
	Email          string `dynamodb:"Email" json:"Email"`
	FirstDecision  bool   `dynamodb:"FirstDecision" json:"FirstDecision"`
	FirstName      string `dynamodb:"FirstName" json:"FirstName"`
	FullName       string `dynamodb:"Fullname" json:"FullName"`
	LastName       string `dynamodb:"LastName" json:"LastName"`
	LastUpdated    int64  `dynamodb:"LastUpdated" json:"LastUpdated"`
	MembershipType string `dynamodb:"MembershipType" json:"MembershipType"`
	NewCreation    bool   `dynamodb:"NewCreation" json:"NewCreation"`
	PhoneNumber    string `dynamodb:"PhoneNumber" json:"PhoneNumber"`
	Rededication   bool   `dynamodb:"Rededication" json:"Rededication"`
	UUID           string `dynamodb:"UUID" json:"UUID"`
	Volunteer      bool   `dynamodb:"Volunteer" json:"Volunteer"`
}

var people []People

func main() {
	// Init Router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/people/{id}", getSingle).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//checks for empty strings in SPOT
func checkValidStr(attr string) string {
	if attr == "" {
		return "n/a"
	}
	return attr
}

//checks for empty Booleans in SPOT
func checkValidBool(attr bool) bool {
	if attr == false {
		return false
	}
	return true
}

// Will pull one person's info
func getSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	// starts session with us-east-1 dynamodb and loads credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	//creates Dynamo Client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("SPOT"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {
				S: aws.String(userID),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//converts string to int64
	createdAt, _ := strconv.ParseInt(*result.Item["CreatedAt"].N, 10, 64)
	lastUpdated, _ := strconv.ParseInt(*result.Item["LastUpdated"].N, 10, 64)

	// TO DO: build a clean struct with the clean data
	// TO DO: return the clean struct in JSON format
	member := People{
		CreatedAt:      createdAt,
		Email:          checkValidStr(*result.Item["Email"].S),
		FirstDecision:  checkValidBool(*result.Item["FirstDecision"].BOOL),
		FirstName:      checkValidStr(*result.Item["FirstName"].S),
		FullName:       checkValidStr(*result.Item["FullName"].S),
		LastName:       checkValidStr(*result.Item["LastName"].S),
		LastUpdated:    lastUpdated,
		MembershipType: checkValidStr(*result.Item["MembershipType"].S),
		NewCreation:    checkValidBool(*result.Item["NewCreation"].BOOL),
		PhoneNumber:    checkValidStr(*result.Item["PhoneNumber"].S),
		Rededication:   checkValidBool(*result.Item["Rededication"].BOOL),
		UUID:           checkValidStr(*result.Item["UUID"].S),
		Volunteer:      checkValidBool(*result.Item["Volunteer"].BOOL),
	}
	fmt.Println(member)

	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	// } //Sprintf formats according to a format specifier and returns the resulting string.

	// if item.UUID == "" {
	// 	fmt.Println("Could not find the UUID specified")
	//}
}
