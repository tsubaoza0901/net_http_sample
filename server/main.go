package main

import (
	"io"
	"log"
	"net/http"
	// ①import package
	// _ "net/http/pprof"
)

func hogeHandler(w http.ResponseWriter, req *http.Request) {
	// HTMLテキストをhttp.ResponseWriterへ書き込む
	io.WriteString(w, `
    <!DOCTYPE html>
    <html lang="ja">
    <head>
      <meta charset="UTF-8">
      <title>Go | net/httpパッケージ</title>
    </head>
    <body>
      <h1>Hello, World!</h1>
      <p>Sample Web Server</p>
    </body>
    </html>
`)
}

func main() {
	http.HandleFunc("/", hogeHandler)

	// // ②add code for pprof
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	log.Println("Start Server")
	http.ListenAndServe(":6060", nil)
}
