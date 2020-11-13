package models

import (
	"database/sql"
)

// exported global variable to access the db
var DB *sql.DB

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func InitDB(dataSourceName string) error {
	var err error

	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	return DB.Ping()
}

func AllBooks() ([]Book, error) {
	rows, err := DB.Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}
	defer DB.Close()

	var bks []Book
	for rows.Next() {
		var bk Book
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
