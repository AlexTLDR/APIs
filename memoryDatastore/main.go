package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Contact struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	Mail string `json:"Mail"`
}

var (
	db  *sql.DB
	err error
)

// var contacts = []contact{
// 	{ID: "1", Name: "Alex B", Mail: "foo@protonmail.com"},
// 	{ID: "2", Name: "Alex Test", Mail: "bar@gmail.com"},
// 	{ID: "3", Name: "Foo", Mail: "foo@gmail.com"},
// }

type server struct {
	contacts []Contact
}

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.01:3308)/testapp")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	/* check db connection */
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connection to db successful")
	}

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
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(s.contacts)
	// w.WriteHeader(http.StatusOK)
	var contacts []Contact

	result, err := db.Query("select * from Contacts")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var contact Contact
		err := result.Scan(&contact.ID, &contact.Name, &contact.Mail)
		if err != nil {
			panic(err.Error())
		}
		contacts = append(contacts, contact)
		json.NewEncoder(w).Encode(contacts)
	}
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
	json.NewEncoder(w).Encode(&Contact{})
	w.WriteHeader(http.StatusOK)
}

func (s *server) createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newContact Contact
	_ = json.NewDecoder(r.Body).Decode(&newContact)
	s.contacts = append(s.contacts, newContact)
	json.NewEncoder(w).Encode(newContact)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, `{"success": "contact created"}`)
}

func (s *server) updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedContact Contact
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
	json.NewEncoder(w).Encode(&Contact{})
	w.WriteHeader(http.StatusOK)
}
