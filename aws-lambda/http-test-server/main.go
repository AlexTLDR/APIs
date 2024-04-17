package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Options struct {
	Professionals           string `json:"professionals"`
	Solvers                 string `json:"solvers"`
	AdditionalProfessionals int    `json:"additionalProfessionals"`
	AdditionalSolvers       int    `json:"additionalSolvers"`
}

type Payload struct {
	Bundle  string  `json:"bundle"`
	UUID    string  `json:"UUID"`
	Name    string  `json:"name"`
	EndDate *string `json:"endDate"`
	Options Options `json:"options"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var p Payload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Received payload: %+v", p)
		log.Printf("Received payload: %+v\n", p) // Print to terminal
	})

	http.ListenAndServe(":8080", nil)
}
