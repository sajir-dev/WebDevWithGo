package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// db, err = sql.Open("mysql", "admin:password@tcp(database-2.cfd81motzhjs.ap-south-1.rds.amazonaws.com:3306)/new_schema01?charset=utf8")
	db, err = sql.Open("mysql", "admin:password@tcp(database-2.cfd81motzhjs.ap-south-1.rds.amazonaws.com:3306)/new_schema01?charset=utf8")
	check(err)
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO users ("james@gmail.com", "password", "james", "50");`)
	check(err)
	_, err = stmt.Exec()
	check(err)

	err = db.Ping()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
package main

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "admin:password@tcp(database-4-go-signup-project.cfd81motzhjs.ap-south-1.rds.amazonaws.com)/users-schema?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}