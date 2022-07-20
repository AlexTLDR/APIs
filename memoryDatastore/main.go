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

// var contacts = []contact{
// 	{ID: "1", Name: "Alex B", Mail: "foo@protonmail.com"},
// 	{ID: "2", Name: "Alex Test", Mail: "bar@gmail.com"},
// 	{ID: "3", Name: "Foo", Mail: "foo@gmail.com"},
// }

type server struct {
	contacts []contact
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	s := server{}
	//Start()
	api.HandleFunc("/contacts", s.getAllContacts).Methods("GET")
	api.HandleFunc("/contacts/{id}", s.getContact).Methods("GET")
	api.HandleFunc("/contacts", s.createContact).Methods("POST")
	api.HandleFunc("/contacts/{id}", s.updateContact).Methods("PUT")
	api.HandleFunc("/contacts/{id}", s.deleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func (s *server) getAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.contacts)
	w.WriteHeader(http.StatusOK)
}

func (s *server) getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, article := range s.contacts {
		if article.ID == params["id"] {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	json.NewEncoder(w).Encode(&contact{})
	w.WriteHeader(http.StatusOK)
}

func (s *server) createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newContact contact
	_ = json.NewDecoder(r.Body).Decode(&newContact)
	s.contacts = append(s.contacts, newContact)
	json.NewEncoder(w).Encode(newContact)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, `{"success": "contact created"}`)
}

func (s *server) updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedContact contact
	_ = json.NewDecoder(r.Body).Decode(&updatedContact)
	params := mux.Vars(r)
	for i, contact := range s.contacts {
		if contact.ID == params["id"] {
			updatedContact.ID = params["id"]
			contact.Name = updatedContact.Name
			contact.Mail = updatedContact.Mail
			s.contacts = append(s.contacts[:i], updatedContact)
			json.NewEncoder(w).Encode(updatedContact)
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprint(w, `{"success": "contact updated"}`)

		}
	}
}

func (s *server) deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, contact := range s.contacts {
		if contact.ID == params["id"] {
			s.contacts = append(s.contacts[:i], s.contacts[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(&contact{})
	w.WriteHeader(http.StatusOK)
}
