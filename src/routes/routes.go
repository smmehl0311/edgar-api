package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"db"
)

func GetToken(w http.ResponseWriter, req *http.Request) {

}

func GetRecentFilings(w http.ResponseWriter, req *http.Request) {

}

func GetLastUpdated(w http.ResponseWriter, req *http.Request) {
	session := db.GetSession()
	json.NewEncoder(w).Encode(db.GetLastUpdated(session))
}

type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var People []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range People {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(People)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	People = append(People, person)
	json.NewEncoder(w).Encode(People)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range People {
		if item.ID == params["id"] {
			People = append(People[:index], People[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(People)
}
