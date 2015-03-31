package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, string(str))
}

func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "%v%v%v", s.Greeting, s.Punct, s.Who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
