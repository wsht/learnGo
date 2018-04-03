package methods

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprint(w, "hello")
}

func WebServers() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprint(w, s)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprint(w, s.Greeting+s.Punct+s.Who)
}

func ExerciseHttpHandlers() {
	http.Handle("/string", String("I'm a frayed knot"))
	http.Handle("/struct", &Struct{"Hello", ":", "Hantong"})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
