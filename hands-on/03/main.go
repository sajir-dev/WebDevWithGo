package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func foo(res http.ResponseWriter, req *http.Request) {
	// parsing template
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("Error parsing template", err)
	}

	// executing template
	err = tpl.ExecuteTemplate(res, "index.gohtml", "Hi from foo")
	if err != nil {
		log.Fatalln("Error executing template", err)
	}
}

func bar(res http.ResponseWriter, req *http.Request) {
	// parsing template
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	// executing template
	err = tpl.ExecuteTemplate(res, "index.gohtml", "Hello from Bar")
	if err != nil {
		log.Fatalln("Error executing template", err)
	}
}

func main() {
	// here in our template we dont have to add base html tags like head and body, the wrapping tags like h1, p, etc.. are enough because the file output will be formatted in that way
	http.Handle("/foo/", http.HandlerFunc(foo))
	http.Handle("/bar/", http.HandlerFunc(bar))

	http.ListenAndServe(":8080", nil)
}
