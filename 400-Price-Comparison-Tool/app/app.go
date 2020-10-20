package app

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

// StartApp ...
func StartApp() {
	router := gin.Default()

	router.GET("/item/:id", controllers.GetItem)
	router.POST("/item/post", controllers.PostItem)

	router.Run(":8080")
}
