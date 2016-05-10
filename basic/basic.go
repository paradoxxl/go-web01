package basic

import (
"fmt"
"html"
"log"
"net/http"
)

func Start() {


http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
fmt.Println("Server started")

}

