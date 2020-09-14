package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// functions runs on starting and stores the templates in a global variable
func init() {
	// template.Must() will check the error
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// 1. os.Stdout prints the parsed gohtml file
	// tpl, err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 2. new file is created here
	// tpl, err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// nf, err := os.Create("index2.html")
	// if err != nil {
	// 	log.Println("Error creating file: ", err)
	// }

	// defer nf.Close()

	// err = tpl.Execute(nf, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 3. Executing templates
	// tpl, err := template.ParseFiles("one.gmao")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Adding two more files to the container
	// tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// executing vespa.gmao from the container
	// err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // executing two.gmao from the container
	// err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // executing one.gmao from the container
	// err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // .Execute simply executes what it gets first. Here in the container one.gmao is first
	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 4. ParseGlob to parse all from the template directory
	// tpl, err := template.ParseGlob("templates/*")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.Execute(os.Stdout, nil)
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// We parse all the templates from a directory using parseGlob then we execute it using Execute template
