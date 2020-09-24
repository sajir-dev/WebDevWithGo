package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/thispartremoved/", http.StripPrefix("/thispartremoved", http.FileServer(http.Dir("abc"))))
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
