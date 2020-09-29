package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func helloworld(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello from AWS")
}
