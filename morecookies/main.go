package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/abundance", abundance)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "some Value",
	})
}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "Some general value",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "Some specific cookie",
	})
}

func read(res http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #1:", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #2:", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(res, "YOUR COOKIE #3:", c3)
	}
}
