package main

import (
	"fmt"
	"log"
	"net/http"

	"bookstore.project/models"
)

func main() {
	// var err error

	// db, err := sql.Open("postgres", "postgres://:password@localhost/book")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := models.InitDB("postgres://:password@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Price, bk.Author)
	}
}
