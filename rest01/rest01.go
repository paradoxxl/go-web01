package rest01

import (
	"log"
	"net/http"
)

func Start() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}