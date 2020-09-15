package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.ParseFiles("tpl-copy.gohtml"))
}

func main() {
	// Passing a variable
	// err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)

	//Passing a string
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-copy.gohtml", `Release self-focus; embrace other-focus`)
	if err != nil {
		log.Fatalln(err)
	}
}
