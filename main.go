package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

//People information in SPOT
//add age group srt, baptized bool, slackID str, ThisIsHome bool, DiscoverYour
type People struct {
	AgeGroup            string `dynamo:"AgeGroup,omitempty" json:"AgeGroup,omitempty"`
	Baptized            bool   `dynamo:"Baptized,omitempty" json:"Baptized,omitempty"`
	Birthday            string `dynamo:"Birthday,omitempty" json:"Birthday,omitempty"`
	Email               string `dynamo:"Email,omitempty" json:"Email,omitemptyv"`
	FirstName           string `dynamo:"FirstName,omitempty" json:"FirstName,omitempty"`
	FullName            string `dynamo:"FullName,omitempty" json:"FullName,omitempty"`
	Gender              string `dynamo:"Gender,omitempty" json:"Gender,omitempty"`
	LastName            string `dynamo:"LastName,omitempty" json:"LastName,omitempty"`
	CreatedAt           int64  `dynamo:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	LastUpdated         int64  `dynamo:"LastUpdated,omitempty" json:"LastUpdated,omitempty"`
	MembershipType      string `dynamo:"MembershipType,omitempty" json:"MembershipType,omitempty"`
	PhoneNumber         string `dynamo:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty"`
	ReceiveEmail        bool   `dynamo:"ReceiveEmail,omitempty" json:"ReceiveEmail,omitempty"`
	SlackID             string `dynamo:"SlackID,omitempty" json:"SlackID,omitempty"`
	UUID                string `dynamo:"UUID,omitempty" json:"UUID,omitempty"`
	NewCreation         bool   `dynamo:"NewCreation,omitempty" json:"NewCreation,omitempty"`
	FirstDecision       bool   `dynamo:"FirstDecision,omitempty" json:"FirstDecision,omitempty"`
	Rededication        bool   `dynamo:"Rededication,omitempty" json:"Rededication,omitempty"`
	Volunteer           bool   `dynamo:"Volunteer,omitempty" json:"Volunteer,omitempty"`
	ThisIsHome          bool   `dynamo:"ThisIsHome,omitempty" json:"ThisIsHome,omitempty"`
	DiscoverYourPurpose bool   `dynamo:"DiscoverYourPurpose,omitempty" json:"DiscoverYourPurpose,omitempty"`
}

var people []People

func main() {
	// Init Router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/people/{id}", getSingle).Methods("GET")
	router.HandleFunc("/api/v1/people", getAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Will pull one person's info
func getSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table("SPOT")

	single := People{}
	err := table.Get("UUID", userID).One(&single)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(&single)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table("SPOT")

	people := []People{}
	err := table.Scan().All(&people)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(&people)
	//fmt.Println(people)
	cleanData := []People{}

	for _, p := range people {
		cleanPerson := People{}
		if p.AgeGroup == "" {
			cleanPerson.AgeGroup = "na"
		} else {
			cleanPerson.AgeGroup = p.AgeGroup
		}
		cleanData = append(cleanData, p)
		return
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.CreatedAt == "" {
			cleanPerson.CreatedAt = "na"
		} else {
			cleanPerson.CreatedAt = p.CreatedAt
		}
		cleanData = append(cleanData, p)
		return
	}

	// loop through people list
	// for each person, create a clean People sruct
	// if field is nil make value "n/a"
	// once clean struct is complete put clean new person in "cleanData"
	// return new lists

}
