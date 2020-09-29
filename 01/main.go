package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cat)
	http.HandleFunc("/mouse", mouse)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<h1>Cat catty</h1><a href="/mouse">mouse</a>`)
}

func mouse(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<h1>Mouse Mousey</h1><a href="/cat">cat</a>`)
}

// parse template for html templates
// pass data to the templates and manipulate them using dot operater, range,  and with other conditional operators
// create a session cookie on logging in
// create a user and a session cookie when sign up
// destroy the session cookie when logout
// set maxage = -1 to destroy the cookie
// while login create a sessionID and match the same with userID for successful login
