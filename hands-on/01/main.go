// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/"
// "/dog/"
// "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

package main

import (
	"io"
	"net/http"
)

func endPoint1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hooyyaaa")
}

func endPoint2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Aiiwaaaa")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Sajir")
}

func main() {
	http.HandleFunc("/", endPoint1)
	http.HandleFunc("/endpoint2/", endPoint2)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
