package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// type user struct {
// 	username string
// 	password string
// 	name     string
// 	age      string
// }

func main() {
	db, err := sql.Open("mysql", "admin:password@tcp(database-4-go-signup-project.cfd81motzhjs.ap-south-1.rds.amazonaws.com)/users-schema?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db.Ping()
	// stmt, _ := db.Prepare(`INSERT INTO users VALUES ("12", "purushu@email.com", "password", "purushu", "42");`)
	// r, _ := stmt.Exec()
	// fmt.Println(r)
	// rows, err := db.Query(`SELECT username, password FROM users;`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()

	// // n, err := db.Exec(stmt)
	// fmt.Println(r)
	// var username, password string
	// for rows.Next() {
	// 	err = rows.Scan(&username, &password)
	// 	// check(err)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("RETRIEVED RECORD:", username, password)
	// }
	username := "james"
	u, err := db.Query(`SELECT password FROM users WHERE username = "` + username + `";`)
	if err != nil {
		fmt.Println("iuhiuhjk", err)
	}
	// fmt.Println(pwd)
	var password string
	u.Next()
	err = u.Scan(&password)
	if err != nil {
		fmt.Println("iuhiuhjk", err)
	}
	fmt.Printf("%v", password)

}
