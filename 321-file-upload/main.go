package main

import (
	"html/template"
	"net/http"
)

type A struct {
	a string `json:"a"`
}

type B struct {
	b1 string `json:"b1"`
	b2 string `json:"b2"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8030", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate("index.gohtml")
}

func upload(w http.ResponseWriter, r *http.Request) {
	a := r.Header
}
