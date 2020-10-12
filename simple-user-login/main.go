package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

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

// var dbSessions = map[string]string{}

// const iota = 10000

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	db, err = sql.Open("mysql", "admin:password@tcp(database-4-go-signup-project.cfd81motzhjs.ap-south-1.rds.amazonaws.com)/users_schema02?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "at index")

}

func signup(w http.ResponseWriter, r *http.Request) {
	// var u user
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		n := r.FormValue("fullname")
		a := r.FormValue("age")

		// userdata := &User{}
		// err := json.NewDecoder(r.Body).Decode(userdata)
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		// fmt.Println(*userdata)
		// fmt.Println(r.Body)
		// fmt.Println(un, p, n, a)
		bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)

		userdata := User{un, bs, n, a}
		fmt.Println(userdata)
		s := `INSERT INTO users03 VALUES ` + `("` + userdata.Username + `", "` + string(userdata.Password) + `", "` + userdata.Name + `", "` + userdata.Age + `");`
		fmt.Println(s)
		stmt, err := db.Prepare(s)
		if err != nil {
			http.Error(w, "Could not connect with server", http.StatusInternalServerError)
			return
		}

		rows, err := stmt.Exec()
		if err != nil {
			http.Error(w, "Username already exists", http.StatusInternalServerError)
			return
		}

		fmt.Println(rows)

		// sID, _:= uuid.New()
		// c := &http.Cookie{
		// 	Name: "session",
		// 	Value: sID.String(),
		// }
		// http.SetCookie(w, c)
		// dbSessions[c.Value] = un

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")

		// var email string
		var password string
		// var name string
		// var age string

		// var u user
		// u, err := db.Query(`SELECT password FROM users WHERE username = "` + username + `";`)
		row, err := db.Query(`SELECT password FROM users03 WHERE username = "` + un + `";`)
		if err != nil {
			fmt.Println("Username and/or password is incorrect")
			return
		}

		// fmt.Printf("%v \n, %T \n", row, row)
		row.Next()
		err = row.Scan(&password)
		if err != nil {
			fmt.Println("Username and/or password is incorrect")
			fmt.Fprintln(w, "Username and/or password is incorrect")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(p))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Wrong credentials")
		}

		fmt.Printf("password: %v", password)
		// sqlStatement := `SELECT * FROM users WHERE username=$1;`
		// row := db.QueryRow(sqlStatement, un)
		// switch err := row.Scan(&email, &password, &name, &age); err {
		// case sql.ErrNoRows:
		// 	fmt.Println("Username and/or password do not match")
		// case nil:
		// 	fmt.Println(email, password, name, age)
		// default:
		// 	panic(err)
		// }

		// if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(p)); err != nil {
		// 	// if two passwords do not match, return a 401 status
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }
		// fmt.Println(email, password, name, age)
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

// getUser(w http.ResponseWriter, r *http.Request){
// 	c, err := r.Cookie("session")
// 	if err != nil {
// 		sID, _ := uuid.New()
// 		c = &http.Cookie {
// 			Name : "session",
// 			Value: sID.String(),
// 		}
// 	}
// 	http.SetCookie(w, c)

// 	var u User

// 	if un, ok := dbSessions[c.Value];ok {
// 		row, err := db.Query(`SELECT password FROM users03 WHERE username = "` + un + `";`)

// 	}
// }
