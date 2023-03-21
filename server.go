package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
