package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each req calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//handler echoes the path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
