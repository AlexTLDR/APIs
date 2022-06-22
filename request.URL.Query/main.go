package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("This is the root path. For the hello, please go to /hello and add the name in the form ?name="))
		fmt.Fprint(w, "This is the root path. For the hello, please go to /hello and add the name in the form ?name=")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		helloFunc(w, r)
	})
	//http.ListenAndServe(":8081", nil)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	log.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)

	panic(http.Serve(listener, nil))

}

func helloFunc(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "world"
	}
	//w.Write([]byte("Hello " + name))
	fmt.Fprint(w, "Hello "+name)
}
