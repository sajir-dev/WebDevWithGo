package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// need a driver to connect to the mysql in aws
	// import driver from packages
	// database/sql package has database manipulation handle
	db, err = sql.Open("mysql", "admin:password@tcp(database-2.cfd81motzhjs.ap-south-1.rds.amazonaws.com:3306)/new_schema01?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
