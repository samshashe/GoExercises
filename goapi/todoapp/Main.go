package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/todos", TodoIndex)
	http.HandleFunc("/todos/{todoid}", TodoShow)
	http.HandleFunc("/todos/create", TodoCreate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
