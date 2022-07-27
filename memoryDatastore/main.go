package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// var (
// 	db  *sql.DB
// 	err error
// )

// var contacts = []contact{
// 	{ID: "1", Name: "Alex B", Mail: "foo@protonmail.com"},
// 	{ID: "2", Name: "Alex Test", Mail: "bar@gmail.com"},
// 	{ID: "3", Name: "Foo", Mail: "foo@gmail.com"},
// }

type server struct {
	contacts []Contact
	db       *sql.DB
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
	s := server{
		db: db,
	}
	//Start()

	api.HandleFunc("/contacts", s.getAllContacts).Methods("GET")
	api.HandleFunc("/contacts/{id}", s.getContact).Methods("GET")
	api.HandleFunc("/contacts", s.createContact).Methods("POST")
	api.HandleFunc("/contacts/{id}", s.updateContact).Methods("PUT")
	api.HandleFunc("/contacts/{id}", s.deleteContact).Methods("DELETE")
	//New functions to add regarding photo manipulation
	//api.HandleFunc("/photos", s.addPhoto).Methods("POST")
	//api.HandleFunc("/photos", s.deletePhoto).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", r))

}

func (s *server) getAllContacts(w http.ResponseWriter, r *http.Request) {

	var contacts []Contact

	result, err := s.db.Query("select * from Contacts")
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
	result, err := s.db.Query("select * from Contacts where ID=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var contact Contact
	for result.Next() {
		err := result.Scan(&contact.ID, &contact.Name, &contact.Mail)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(contact)

}

func (s *server) createContact(w http.ResponseWriter, r *http.Request) {
	query, err := s.db.Prepare("insert into Contacts (ID, Name, Mail) VALUES(?,?,?) ")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	ID := keyVal["ID"]
	Name := keyVal["Name"]
	Mail := keyVal["Mail"]
	_, err = query.Exec(ID, Name, Mail)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprint(w, `{"success": "contact created"}`)
}

func (s *server) updateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query, err := s.db.Prepare("update Contacts set name = ? where ID =?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["Name"]
	_, err = query.Exec(newName, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Contact with ID = %s was updated", params["id"])
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
