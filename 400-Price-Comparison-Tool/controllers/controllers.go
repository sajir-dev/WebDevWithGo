package controllers

import (
	"encoding/json"
	"net/http"

	"../services"
	"github.com/gin-gonic/gin"
)

// GetItem ...
func GetItem(c *gin.Context) {
	id := c.Param("id")
	// fmt.Println(id)
	// c.String(http.StatusOK, "Item %v is returned", id)
	item, _ := json.Marshal(services.GetItem(id))
	c.String(http.StatusOK, string(item))
}

// PostItem ...
func PostItem(c *gin.Context) {
	// id := c.Query("id")
	// page := c.DefaultQuery("page", "0")
	// name := c.PostForm("name")
	// message := c.PostForm("message")

	// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	// id := c.Param("id")
	// // fmt.Println(id)
	// // c.String(http.StatusOK, "Item %v is returned", id)
	// item, _ := json.Marshal(services.GetItem(id))
	// c.String(http.StatusOK, string(item))

	itemid := "c301"
	category := "mobile cover"
	var price float64 = 360.50
	var rating float32 = 4.5

	// c.JSON(200, gin.H{
	// 	"status":"posted",
	// 	"message":""
	// })

	status := services.PostItem(itemid, category, price, rating)
	if status {
		c.String(http.StatusOK, "item added successfully")
		return
	}
	c.String(http.StatusNotFound, "operation could not be completed")
}
