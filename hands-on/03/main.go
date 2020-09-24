package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("public"))))
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
