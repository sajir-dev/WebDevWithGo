package main

import (
	"io"
	"net/http"
)

// type hotdog int

// func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "dog dogg dogggy")
// }
func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dogg doggy")
}

// type hotcat int

// func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "cat cat catty")
// }

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat catt catty")
}

func main() {
	// var d hotdog
	// var c hotcat

	// mux := http.NewServeMux()
	// mux.Handle("/dog/", d) // run for anything comes after dog/
	// mux.Handle("/cat", c)  // run for exact match
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	// http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":8080", nil)
}
