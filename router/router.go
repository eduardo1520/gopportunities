package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the Gin router
	router := gin.Default()
	// Define the route and handler
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Start the server
	router.Run(":8080")
}
