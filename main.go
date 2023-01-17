package main

import (
	"go-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new gin instance
	r := gin.Default()

	// Load the routes
	routes.AuthRoutes(r)

	// Run the server
	r.Run(":8080")
}
