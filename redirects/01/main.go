package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo: ", req.Method)
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)
	// res.Header().Set("Location", "/")
	// res.WriteHeader(http.StatusSeeOther)
	// http.Redirect(res, req, "/", http.StatusSeeOther)
	// http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
