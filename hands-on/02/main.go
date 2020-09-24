package main

import (
	"log"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

// func foo(res http.ResponseWriter, req *http.Request) {

// }
