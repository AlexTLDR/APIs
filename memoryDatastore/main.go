package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type contact struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	Mail string `json:"Mail"`
}

var contacts = []contact{
	{ID: "1", Name: "Alex B", Mail: "foo@protonmail.com"},
	{ID: "2", Name: "Alex Test", Mail: "bar@gmail.com"},
	{ID: "3", Name: "Foo", Mail: "foo@gmail.com"},
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/contacts", getAllContacts).Methods("GET")
	api.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	api.HandleFunc("/contacts", createContact).Methods("POST")
	api.HandleFunc("/contacts/{id}", updateContact).Methods("PUT")
	api.HandleFunc("/contacts/{id}", deleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func getAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
	w.WriteHeader(http.StatusOK)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, article := range contacts {
		if article.ID == params["id"] {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode(&contact{})
	w.WriteHeader(http.StatusOK)
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newContact contact
	_ = json.NewDecoder(r.Body).Decode(&newContact)
	contacts = append(contacts, newContact)
	json.NewEncoder(w).Encode(newContact)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, `{"success": "contact created"}`)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedContact contact
	_ = json.NewDecoder(r.Body).Decode(&updatedContact)
	params := mux.Vars(r)
	for i, contact := range contacts {
		if contact.ID == params["id"] {
			updatedContact.ID = params["id"]
			contact.Name = updatedContact.Name
			contact.Mail = updatedContact.Mail
			contacts = append(contacts[:i], updatedContact)
			json.NewEncoder(w).Encode(updatedContact)
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprint(w, `{"success": "contact updated"}`)

		}
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, contact := range contacts {
		if contact.ID == params["id"] {
			contacts = append(contacts[:i], contacts[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(&contact{})
	w.WriteHeader(http.StatusOK)
}
