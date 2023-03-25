# golang-sample-web-server

## Development Environment

- MacBook Pro(M1,2021)
- macOS Monterey v12.2.1 
- go version go1.20 darwin/arm64

## Directory Structure

```shell
.
├── .env
├── .gitignore
├── README.md
├── go.mod
├── index.html
├── main.go
├── scan.go
├── server.go
└── ward.txt

```

## Build

- Start Web Server

```shell
% go build .
% ./main
```

## Request

```shell
% curl 'http://localhost:8080/'
// Hello World!!
```

## Reference

- [GoでWebサーバーを構築](https://zenn.dev/satumahayato010/articles/b4a6cccbb3bb09)
- [テンプレート機能を使用する (text/template, html/template)](https://maku77.github.io/p/z8behko/)
- [Go でファイルや標準入力からテキストを一行ずつ読む](https://www.yunabe.jp/tips/golang_readlines.html)