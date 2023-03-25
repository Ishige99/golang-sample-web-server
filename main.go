package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/create", CreateWardHandler)

	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
