package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, x := c.FormFile("file")
		username, y := c.GetPostForm("username")
		password, z := c.GetPostForm("password")
		fmt.Println(username, password)
		fmt.Println(x, y, z)

		// Upload file to a designated place
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s uploaded", file.Filename))
	})
	router.Run(":8080")
}
