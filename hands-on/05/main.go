package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
