package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to ThrottlePoint")
	r := gin.Default()

	// Home route for sanity check
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to ThrottlePoint",
		})
	})

	// Start the server
	r.Run("0.0.0.0:8080")
}
