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
	Id        string `json:"id"`
	User_name string `json:"user_name"`
	Mail      string `json:"mail"`
}

type server struct {
	contacts []Contact
	db       *sql.DB
}

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.01:3306)/aiggato")
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

	result, err := s.db.Query("SELECT id, user_name, mail FROM contacts;")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var contact Contact
		err := result.Scan(&contact.Id, &contact.User_name, &contact.Mail)
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
	result, err := s.db.Query("SELECT id, user_name, mail FROM contacts where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var contact Contact
	for result.Next() {
		err := result.Scan(&contact.Id, &contact.User_name, &contact.Mail)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(contact)

}

func (s *server) createContact(w http.ResponseWriter, r *http.Request) {
	query, err := s.db.Prepare("insert into contacts (id, user_name, mail) VALUES(?,?,?) ")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]
	user_name := keyVal["user_name"]
	mail := keyVal["mail"]
	_, err = query.Exec(id, user_name, mail)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprint(w, `{"success": "contact created"}`)
}

func (s *server) updateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query, err := s.db.Prepare("update contacts set mail = ?, user_name = ? where id =?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newMail := keyVal["mail"]
	newName := keyVal["user_name"]
	_, err = query.Exec(newMail, newName, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "contact with id = %s was updated", params["id"])
}

func (s *server) deleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	query, err := s.db.Prepare("delete from contacts where id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "contact with id = %s was deleted", params["id"])
}
