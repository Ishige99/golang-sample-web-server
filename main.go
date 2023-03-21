package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
