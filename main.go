package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	route := app.Group("/api/v1")

	route.GET("/items", getItemsHandler)

	app.Run(":8080")
}

func getItemsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"items": "Hello World",
	})
}