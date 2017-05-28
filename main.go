package main

import (
	"log"
	"net/http"
	//"time"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
