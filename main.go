package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	ID          string `json:"ID"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	DateOfBirth string `json:"DateOfBirth"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"Email"`
}
type allUsers []user

var users = allUsers{
	{
		ID:          "1",
		FirstName:   "Man",
		LastName:    "Creating http json Api in go",
		DateOfBirth: "02/06/2002",
		PhoneNumber: "493753873",
		Email:       "sfiuofsdsklj",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome  home")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data in order")
	}

	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}
func getOneUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	for _, singleuser := range users {
		if singleuser.ID == userID {
			json.NewEncoder(w).Encode(singleuser)
		}
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var updatedEvent user

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range users {
		if singleEvent.ID == userID {
			singleEvent.FirstName = updatedEvent.FirstName
			singleEvent.LastName = updatedEvent.LastName
			singleEvent.DateOfBirth = updatedEvent.DateOfBirth
			singleEvent.PhoneNumber = updatedEvent.PhoneNumber
			singleEvent.Email = updatedEvent.Email

			users = append(users[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)

		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars((r))["id"]

	for i, singleEvent := range users {
		if singleEvent.ID == eventID {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)

		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createUser).Methods("POST")
	router.HandleFunc("/events", getAllUsers).Methods("GET")
	router.HandleFunc("/events/{id}", getOneUser).Methods("GET")
	router.HandleFunc("/events/{id}", updateUser).Methods(("PATCH"))
	router.HandleFunc("/events/{id}", deleteUser).Methods(("DELETE"))
	log.Fatal(http.ListenAndServe(":8080", router))

}
