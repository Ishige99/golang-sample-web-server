package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//
// テスト用のハンドラ："Hello World!!"を返すのみ
//

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hit: /")

	response := []byte("Hello World!!")
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

//
// index.html用のハンドラ：index.htmlの内容を返します
//

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hit: /index")

	wardList := FileRead("ward.txt")
	index, err := template.ParseFiles("index.html") // htmlテンプレートファイル(index.html)の解析を行います。

	if err != nil {
		log.Fatal(err)
	}

	getWards := New(wardList)

	// 解析したhtmlテンプレートファイルを実行してレスポンスに出力
	if err := index.Execute(w, getWards); err != nil {
		log.Fatal(err)
	}
}

//
// 投稿用のハンドラ：区を追加します
//

func CreateWardHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	formValue := req.FormValue("value")
	file, err := os.OpenFile("ward.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(file, formValue)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, req, "/index", http.StatusFound)
}
