package main

import (
	"html/template"
	"log"
	"os"
)

var tplcopy *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tplcopy = template.Must(template.ParseFiles("tplcopy.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	// err := tpl.Execute(os.Stdout, buddha)
	err := tplcopy.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
