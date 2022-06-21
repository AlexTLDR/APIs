package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//handlerFunc(w, r)
		w.Write([]byte("This is the root path. For the hello, please go to /hello and add the name in the form ?name="))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		helloFunc(w, r)
	})
	log.Println("The port on which the service is running is 8081")
	http.ListenAndServe(":8081", nil)

}

//Isn't it simpler to use fmt.Fprint instead of w.Write([]byte("my string"))?

/*func handlerFunc(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		fmt.Fprint(w, "This is the root path. For the hello, please go to /hello")
	} else if r.URL.Path == "/hello" {
		fmt.Fprint(w, "Hello world!")
	}
}*/

func helloFunc(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	if r.FormValue("name") == "" {
		name = "world"
	}
	w.Write([]byte("Hello " + name))
}
