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

// // Will pull everyone's info on database
func getAll(w http.ResponseWriter, r *http.Request) {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table("SPOT")

	people := []People{}
	err := table.Scan().All(&people)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(&people)

	// loop through all data and check for empty fields, clean up data and append to struct
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
		if p.Baptized == false {
			cleanPerson.Baptized = false
		} else {
			cleanPerson.Baptized = p.Baptized
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.Birthday == "" {
			cleanPerson.Birthday = "na"
		} else {
			cleanPerson.Birthday = p.Birthday
		}
		cleanData = append(cleanData, p)
		return
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.Email == "" {
			cleanPerson.Email = "na"
		} else {
			cleanPerson.Email = p.Email
		}
		cleanData = append(cleanData, p)
		return
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.FirstName == "" {
			cleanPerson.FirstName = "na"
		} else {
			cleanPerson.FirstName = p.FirstName
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.FullName == "" {
			cleanPerson.FullName = "na"
		} else {
			cleanPerson.FullName = p.FullName
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.Gender == "" {
			cleanPerson.Gender = "na"
		} else {
			cleanPerson.Gender = p.Gender
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.LastName == "" {
			cleanPerson.LastName = "na"
		} else {
			cleanPerson.LastName = p.LastName
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.CreatedAt != 0 {
			cleanPerson.CreatedAt = 0
		} else {
			cleanPerson.CreatedAt = p.CreatedAt
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.LastUpdated == 0 {
			cleanPerson.LastUpdated = 0
		} else {
			cleanPerson.LastUpdated = p.LastUpdated
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.MembershipType == "" {
			cleanPerson.MembershipType = "na"
		} else {
			cleanPerson.MembershipType = p.MembershipType
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.PhoneNumber == "" {
			cleanPerson.PhoneNumber = "na"
		} else {
			cleanPerson.PhoneNumber = p.PhoneNumber
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.ReceiveEmail == false {
			cleanPerson.ReceiveEmail = false
		} else {
			cleanPerson.ReceiveEmail = p.ReceiveEmail
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.SlackID == "" {
			cleanPerson.SlackID = "na"
		} else {
			cleanPerson.SlackID = p.SlackID
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.UUID == "" {
			cleanPerson.UUID = "na"
		} else {
			cleanPerson.UUID = p.UUID
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.NewCreation == false {
			cleanPerson.NewCreation = false
		} else {
			cleanPerson.NewCreation = p.NewCreation
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.FirstDecision == false {
			cleanPerson.FirstDecision = false
		} else {
			cleanPerson.FirstDecision = p.FirstDecision
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.Rededication == false {
			cleanPerson.Rededication = false
		} else {
			cleanPerson.Rededication = p.Rededication
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.Volunteer == false {
			cleanPerson.Volunteer = false
		} else {
			cleanPerson.Volunteer = p.Volunteer
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.ThisIsHome == false {
			cleanPerson.ThisIsHome = false
		} else {
			cleanPerson.ThisIsHome = p.ThisIsHome
		}
		cleanData = append(cleanData, p)
	}

	for _, p := range people {
		cleanPerson := People{}
		if p.DiscoverYourPurpose == false {
			cleanPerson.DiscoverYourPurpose = false
		} else {
			cleanPerson.DiscoverYourPurpose = p.DiscoverYourPurpose
		}
		cleanData = append(cleanData, p)
	}
}
