package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the root path. For the hello, please go to /hello and add the name in the form ?name="))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		helloFunc(w, r)
	})
	log.Println("The port on which the service is running is 8081")
	http.ListenAndServe(":8081", nil)

}

func helloFunc(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "world"
	}
	w.Write([]byte("Hello " + name))
}
