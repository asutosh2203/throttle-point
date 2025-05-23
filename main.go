package main

import (
	"fmt"

	"github.com/asutosh2203/throttle-point.git/handlers"
	"github.com/asutosh2203/throttle-point.git/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to ThrottlePoint")
	r := gin.Default()

	r.Any("/*proxyPath", middleware.RateLimiter(), handlers.ProxyHandler)

	// Start the server
	r.Run("0.0.0.0:8080")
}
