package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/index", IndexHandler)

	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
