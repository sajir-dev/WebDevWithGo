package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Anything that you want in this function")
}

func main() {
	var hd hotdog
	http.ListenAndServe(":8080", hd) // ListenAndServe has first param a route string and second a handler. A handler is nothing but that implements ServeHTTP method. ServeHTTP takes in a responseWriter and a pointer to an interface
}
