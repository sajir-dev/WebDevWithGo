package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	c = appendFileNames(res, req)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(res, "index.gohtml", xs)
}

func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		return c
	}
	return c
}

func appendFileNames(res http.ResponseWriter, req *http.Request) *http.Cookie {
	c, _ := req.Cookie("session")
	p1 := "sunset.jpg"
	p2 := "beach.jpg"
	p3 := "night.jpg"

	s := c.Value

	if !strings.Contains(s, p1) {
		s += "|" + p1
	}

	if !strings.Contains(s, p2) {
		s += "|" + p2
	}

	if !strings.Contains(s, p3) {
		s += "|" + p3
	}

	c.Value = s
	http.SetCookie(res, c)
	return c
}
