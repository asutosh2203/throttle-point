package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var buckets = make(map[string]*TokenBucket)
var bucketsMutex sync.Mutex

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()

		bucketsMutex.Lock()
		bucket, exists := buckets[ip]
		if !exists {
			bucket = NewTokenBucket(10, 6) // 10 tokens max, refill 1/6sec
			buckets[ip] = bucket
		}
		bucketsMutex.Unlock()

		if !bucket.AllowRequest() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Slow down, cowboy!"})
			return
		}

		ctx.Next()
	}
}
