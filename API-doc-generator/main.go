package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"documentation-generator/apis"

	"github.com/gorilla/mux"
)

const port = 3000

type Func struct {
	Name    string
	Comment string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("ui/html/get.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(content))
	})

	r.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("ui/html/post.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(content))
	})
	r.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("ui/html/put.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(content))
	})
	r.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("ui/html/delete.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(content))
	})
	r.HandleFunc("/footer.html", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("ui/html/footer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(content))
	})
	r.HandleFunc("/users/all", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apis.Users)
	})

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		apis.GetUserByID(w, r)
	}).Methods("GET")
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		apis.CreateUser(w, r)
	}).Methods("POST")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		apis.EditUser(w, r)
	}).Methods("PUT")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		apis.DeleteUser(w, r)
	}).Methods("DELETE")

	fmt.Printf("Server is running on port %d\n Link to documentation -> http://localhost:6060/pkg/documentation-generator/apis", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
