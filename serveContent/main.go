package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/toby.jpg">`)
	fmt.Println(res)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(res, "Error opening file", 404)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
	fmt.Println(res)
}
