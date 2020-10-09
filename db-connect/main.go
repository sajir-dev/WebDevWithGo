package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "admin:password@tcp(database-4-go-signup-project.cfd81motzhjs.ap-south-1.rds.amazonaws.com)/users-schema?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db.Ping()
	stmt, _ := db.Prepare(`INSERT INTO users VALUES ("12", "james@email.com", "password", "james", "42");`)
	r, _ := stmt.Exec()
	fmt.Println(r)
	// // rows, err := db.Query(`SELECT password FROM users;`)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // defer db.Close()

	// // n, err := db.Exec(stmt)
	// fmt.Println(r)
	// // var name string
	// for r.Next() {
	// 	err = r.Scan(&name)
	// 	// check(err)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("RETRIEVED RECORD:", name)
	// }
}
