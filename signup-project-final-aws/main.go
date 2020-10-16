package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	Username string
	Password []byte
	Name     string
	Age      string
}

var tpl *template.Template

var db *sql.DB
var err error

var dbSessions = map[string]string{}

var userdata User

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	db, err = sql.Open("mysql", "admin:Admin@123@tcp(0.0.0.1:3306)/usersdb")
	if err != nil {
		log.Fatalln("could not establish database connection", err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", userdata)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		pw := r.FormValue("password")
		name := r.FormValue("name")
		age := r.FormValue("age")

		bs, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)

		s := `INSERT INTO users VALUES ("` + un + `","` + string(bs) + `","` + name + `","` + age + `");`
		stmt, err := db.Prepare(s)
		// fmt.Println(s)
		if err != nil {
			http.Error(w, "Could not connect with server", http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Username already exists", http.StatusInternalServerError)
			return
		}

		userdata = User{un, bs, name, age}

		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", userdata)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		pw := r.FormValue("password")

		fmt.Println(un, pw)

		row := db.QueryRow(`SELECT * FROM users where username= "` + un + `";`)
		fmt.Println(row)
		err := row.Scan(&userdata.Username, &userdata.Password, &userdata.Name, &userdata.Age)
		fmt.Println(userdata)
		if err != nil {
			http.Error(w, "wrong credentials", http.StatusNotFound)
			return
		}

		// fmt.Println("password from db: ", userdata.Password)
		// bss, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		// fmt.Println("password typed: ", bss)
		err = bcrypt.CompareHashAndPassword([]byte(userdata.Password), []byte(pw))
		if err != nil {
			http.Error(w, "wrong credentials", http.StatusNotFound)
			return
		}

		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		dbSessions[c.Value] = un

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "login.gohtml", userdata)
}

func bar(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	age, _ := strconv.ParseInt(userdata.Age, 10, 32)
	if age < 18 {
		fmt.Fprintln(w, "You are not allowed in the bar")
	} else {
		fmt.Fprintln(w, "Take your drink")
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	var u User
	userdata = u

	c, _ := r.Cookie("session")

	delete(dbSessions, c.Value)

	sID := uuid.New()
	c = &http.Cookie{
		Name:   "session",
		Value:  sID.String(),
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	fmt.Println(un)
	row := db.QueryRow(`SELECT * FROM users where username= "` + un + `";`)
	fmt.Println(row)
	err = row.Scan(&userdata.Username, &userdata.Password, &userdata.Name, &userdata.Age)
	fmt.Println(userdata.Username)
	if userdata.Username == "" {
		// fmt.Println("it is nil")
		return false
	}
	return true
}
