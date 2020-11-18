package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

type Item struct {
	ItemID      string `json:"itemid"`
	ItemName    string `json:"item"`
	Price       string `json:"price"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	MarketPlace string `json:"market_place"`
}

func init() {
	DB, err = sql.Open("postgres", "postgres://postgres:password@localhost/postgres?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("You are connected to local postgres")
}

func main() {
	rows, err := DB.Query("select * from items join marketplace using(itemid) where marketplace = 'amazon';")
	// rows, err := DB.Query(`select * from items;`)
	fmt.Println(rows)
	if err != nil {
		fmt.Println("query err", err)
		panic(err)
	}

	var Items []Item
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Description, &item.Rating, &item.MarketPlace)
		if err != nil {
			fmt.Println("scan err", err)
			panic(err)
		}
		Items = append(Items, item)
	}
	fmt.Println(Items)
}
