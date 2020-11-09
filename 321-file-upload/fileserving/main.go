package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use()
	router.Static("/assets1", "assets")
	router.Static("/assets2", "/assets")
	router.StaticFS("/more_static", gin.Dir("/my_files", true))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.GET("/test", func(c *gin.Context) {
		c.FileFromFS("/my_files", gin.Dir("/my_files", false))
	})
	// router.Use(static.Serve("/", static.LocalFile("/assets", false)))
	router.GET("/", func(c *gin.Context) {
		c.String(200, "I'm running")
	})

	router.Run(":8080")
}
