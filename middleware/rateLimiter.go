package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var buckets = make(map[string]*TokenBucket)
var bucketsMutex sync.Mutex

const cleanupInterval = 5 * time.Minute
const bucketTTL = 10 * time.Minute

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.RemoteIP()

		bucketsMutex.Lock()
		bucket, exists := buckets[ip]
		if !exists {
			bucket = NewTokenBucket(10, 6) // 10 tokens max, refill 1/6sec
			buckets[ip] = bucket
		}
		bucketsMutex.Unlock()

		if !bucket.AllowRequest() {
			log.Printf("Rate limit exceeded for IP: %s", ip)
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Slow down, cowboy!"})
			return
		}

		ctx.Next()
	}
}

// clean up token buckets that are not used for 10 minutes
func StartBucketCleanup() {
	go func() {
		for {
			time.Sleep(cleanupInterval)

			bucketsMutex.Lock()
			now := time.Now()

			for ip, bucket := range buckets {
				bucket.Mutex.Lock()

				if now.Sub(bucket.LastRefillTime) > bucketTTL {
					delete(buckets, ip)
				}
				bucket.Mutex.Unlock()
			}

			bucketsMutex.Unlock()

		}
	}()
}
