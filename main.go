package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)

	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
