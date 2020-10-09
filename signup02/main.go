package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// var db *sql.DB
// var err error

type user struct {
	Username string
	Password []byte
	Name     string
	Age      string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	// db, err = sql.Open("mysql", "admin:password@tcp(database-2.cfd81motzhjs.ap-south-1.rds.amazonaws.com:3306)/new_schema01?charset=utf8")
	// if err != nil {
	// 	fmt.Println("Could not establish database connection")
	// 	log.Fatalln(err)
	// 	panic(err)
	// } else {
	// 	fmt.Println("db connection success")
	// }

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u user
	// catching form
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("name")
		a := r.FormValue("age")
		fmt.Println(un, p, f, a)
		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, a}
		dbUsers[un] = u

		// stmt, err := db.Prepare(`INSERT INTO users VALUES (" ` + un + `,` + string(bs) + `,` + f + `,` + a + `");`)
		// fmt.Println(stmt)
		// if err != nil {
		// 	http.Error(w, "could not create user", http.StatusInternalServerError)
		// }
		// defer stmt.Close()

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", u)

}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u user
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		// is there a usename?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func bar(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	u := getUser(w, r)
	age, _ := strconv.Atoi(dbUsers[u.Username].Age)
	if age < 18 {
		fmt.Fprintln(w, "You are not allowed into the bar")
	} else {
		fmt.Fprintln(w, "Come take your virtual drinks")
	}
}

func getUser(w http.ResponseWriter, r *http.Request) user {
	var u user
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)

	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	fmt.Println(u)
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		// dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	// c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
