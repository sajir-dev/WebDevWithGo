package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	Username string
	Password []byte
	Name     string
	Age      string
}

var userdata User

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3309)/testusers")
	if err != nil {
		fmt.Println("Not connected to mysql local", err)
		return
	}
	defer db.Close()

	fmt.Println("connected successfully")

	// _, err = db.Query(`INSERT INTO users VALUES ( "demo", "demopwd", "demoname", "45" )`)

	fmt.Println(userdata)

	row := db.QueryRow(`SELECT * FROM users WHERE username="test@email.com"`)
	if err != nil {
		fmt.Println("could not perform the operation", err)
	}
	row.Scan(&userdata.Username, &userdata.Password, &userdata.Name, &userdata.Age)
	fmt.Println(userdata)
	// fmt.Println(stmt)
}
