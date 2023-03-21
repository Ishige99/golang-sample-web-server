package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hit: /")

	response := []byte("Hello World!!")
	_, err := w.Write(response)
	if err != nil {
		// エラー発生時にサーバーが停止しないように、Printlnにしています。
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
