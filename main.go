package main

import (
	"parapilot/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register routes
	router.GET("/passages", handlers.PassagesHandler)

	router.Run(":8000")
}
