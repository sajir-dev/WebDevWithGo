package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/"))
}

func main() {
	db, err = sql.Open("mysql", "admin:password@tcp(database-4-go-signup-project.cfd81motzhjs.ap-south-1.rds.amazonaws.com)/users-schema?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

}

func signup(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {

}

func bar(w http.ResponseWriter, r *http.Request) {

}

func logout(w http.ResponseWriter, r *http.Request) {

}

func getUser(un string) user {
	var u user
	pwd, err := db.Query(`SELECT * FROM users WHERE username = "` + username + `";`))
	if err != nil {
		fmt.Println("error getting password", err)
		return user
	}
	return user
}