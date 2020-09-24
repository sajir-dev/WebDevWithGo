package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("index.gohtml"))
	err := tpl.ExecuteTemplate(res, "index.gohtml", "Hello from dog")
	if err != nil {
		log.Println(err)
	}
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}
