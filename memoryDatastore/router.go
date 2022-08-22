package main

import (
	"github.com/gorilla/mux"
)

func router(s server) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	//Start()

	api.HandleFunc("/contacts", s.getAllContacts).Methods("GET")
	api.HandleFunc("/contacts/{id}", s.getContact).Methods("GET")
	api.HandleFunc("/contacts", s.createContact).Methods("POST")
	api.HandleFunc("/contacts/{id}", s.updateContact).Methods("PUT")
	api.HandleFunc("/contacts/{id}", s.deleteContact).Methods("DELETE")
	//New functions to add regarding photo manipulation
	//api.HandleFunc("/photos", s.addPhoto).Methods("POST")
	//api.HandleFunc("/photos", s.deletePhoto).Methods("DELETE")
	return r
}
