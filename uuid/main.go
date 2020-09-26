package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}

		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Fprintf(res, `cookie: %v`, cookie)
	fmt.Println(cookie)
}
