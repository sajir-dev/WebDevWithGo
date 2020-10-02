package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name  string
	Age   int
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", marsh)
	http.HandleFunc("/encd", encod)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FOO</title>
	</head>
	<body>
	You are at foo
	</body>
	</html>`
	res.Write([]byte(s))
}

func marsh(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	p1 := person{
		Name:  "James",
		Age:   45,
		Items: []string{"suit", "gun", "Wry sense of humor"},
	}
	// marshal to a variable j
	j, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}

	res.Write([]byte(j))
}

func encod(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	p1 := person{
		Name:  "Miss moneypenny",
		Age:   33,
		Items: []string{"knife", "skirt", "lipstick"},
	}
	// send response directly
	err := json.NewEncoder(res).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
