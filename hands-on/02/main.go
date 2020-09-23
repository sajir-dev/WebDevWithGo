// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func foo(res http.ResponseWriter, req *http.Request) {
	// io.WriteString(res, "Hi from foo")
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// tpl.Execute(os.Stdout, "Hi from foo")
	tpl.ExecuteTemplate(res, "index.gohtml", "Hi from foo")
}

func bar(res http.ResponseWriter, req *http.Request) {
	// io.WriteString(res, "Hello from bar")
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// tpl.Execute(os.Stdout, "Hello from bar")
	tpl.ExecuteTemplate(res, "index.gohtml", "Hello from bar")
}

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar/", bar)
	// http.Handle("/foo", http.HandlerFunc(foo)) // conversion to use handle function
	// http.Handle("/bar/", http.HandlerFunc(bar)) //conversion to use handle function

	http.ListenAndServe(":8080", nil)
}
