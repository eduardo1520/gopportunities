package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the Gin router
	router := gin.Default()
	// Define the route and handler
	InitializeRoutes(router)
	// Start the server
	router.Run(":8080")
}
