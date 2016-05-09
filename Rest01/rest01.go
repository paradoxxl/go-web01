package rest01

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/blubb",BlubbIndex)
	router.HandleFunc("/blubb/{blubbId}",BlubbShow)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func BlubbIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Blubb Index, %q", html.EscapeString(r.URL.Path))
}

func BlubbShow(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	Id := vars["blubbId"]
	fmt.Fprintf(w, "Blubb: , %q", Id)
}
