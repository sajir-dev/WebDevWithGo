package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe(":8080", r)
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
	</body>
	</html>
	`
	res.Header().Set("Content-Type", "text/html; charset: utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(s))
}

func getUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "M",
		Age:    45,
		ID:     "1001",
	}

	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	// Write content-type, statuscode, payload
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(res, "%s\n", uj)
}
