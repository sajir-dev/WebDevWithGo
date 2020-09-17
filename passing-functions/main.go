package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

// template.FuncMap is a Map that maps functions with keys so that we can use them in the templates
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Mohammaed",
		Motto: "Inner soul",
	}

	sages := []sage{b, g, m}

	// Execute and ExecuteTemplates are completely different and dont fall for that error again
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
