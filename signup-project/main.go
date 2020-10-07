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
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	// fmt.Fprint(w, "at index")
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "at signup")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "at login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "at logout")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
