package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) getAllContacts(w http.ResponseWriter, r *http.Request) {

	var contacts []Contact

	result, err := s.db.Query("SELECT id, user_name, mail FROM contacts;")
	if err != nil {
		http.Error(w, "can't delete id", 500)
		fmt.Fprintf(w, "add error here\n")
		return
	}

	defer result.Close()

	for result.Next() {
		var contact Contact
		err := result.Scan(&contact.Id, &contact.User_name, &contact.Mail)
		if err != nil {
			panic(err.Error())
		}
		contacts = append(contacts, contact)
		log.Println(contacts)
	}
	json.NewEncoder(w).Encode(contacts)
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

	body, err := io.ReadAll(r.Body)
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

	body, err := io.ReadAll(r.Body)
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

	result, err := query.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected != 1 {
		http.Error(w, "can't delete id", 500)
		fmt.Fprintf(w, "could not delete contact with id = %s\n", params["id"])
		return
	}

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "contact with id = %s was deleted", params["id"])
}
