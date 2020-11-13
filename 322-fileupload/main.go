package main

import (
	"github.com/gin-gonic/gin"
)

// func main() {
// 	http.Handle("/", http.FileServer(http.Dir("./static")))
// 	http.ListenAndServe(":3000", nil)
// }

func main() {
	r := gin.New()
	r.Static("/images/", "./static")
	r.Run(":3000")
}
