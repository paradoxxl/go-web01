package rping

import (
	"net/http"
	"log"
)

func Start(){

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
