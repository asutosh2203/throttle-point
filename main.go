package main

import (
	"fmt"
	"log"

	"github.com/asutosh2203/throttle-point.git/handlers"
	"github.com/asutosh2203/throttle-point.git/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to ThrottlePoint")

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	r := gin.Default()

	r.Any("/*proxyPath", middleware.RateLimiter(), handlers.ProxyHandler)

	// Start the server
	r.Run("0.0.0.0:8080")
}
