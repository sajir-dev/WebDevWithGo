package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// name := "Sajir"
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	tpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h2>` + name + `</h2>
		</body>
		</html>
	`
	fmt.Println(tpl)

	// 1. go run main.go > index1.html to yield the html file

	// 2. io.Copy(nf, strings.NewReader will also create a file)

	// 3. os.Args[1] command line input along with go run main.go os.Args[0] is the running binary itself
	nf, err := os.Create("index3.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
