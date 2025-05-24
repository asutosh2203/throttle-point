package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/asutosh2203/throttle-point.git/ai"
	"github.com/gin-gonic/gin"
)

// construct the predictor
var predictor = ai.NewRuleBasedPredictor()

var buckets = make(map[string]*TokenBucket)
var bucketsMutex sync.Mutex

const cleanupInterval = 5 * time.Minute
const bucketTTL = 10 * time.Minute

const updateThreshold = 10

// 1 -> 3 seconds => 1 seconds -> 1/3 tokens

func getRefillRate(risk float64) float64 {
	switch {
	case risk <= 0.3:
		return float64(1) / float64(3)
	case risk <= 0.7:
		return float64(1) / float64(10)
	case risk <= 0.85:
		return float64(1) / float64(20)
	default:
		return float64(1) / float64(60)
	}
}

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.RemoteIP()

		intent, err := predictor.PredictIntent(ctx.Request)

		if err != nil {
			log.Default().Println(err)
		}

		bucketsMutex.Lock()
		bucket, exists := buckets[ip]

		// create token bucket if it doesn't exist
		if !exists {

			refillRate := getRefillRate(intent.RiskScore)

			bucket = NewTokenBucket(10, refillRate)
			buckets[ip] = bucket
		} else {
			bucket.RequestCount++

			// update the token refill rate based on recalculated intent
			if bucket.RequestCount > updateThreshold {
				bucket.RequestCount = 0

				newIntent, err := predictor.PredictIntent(ctx.Request)

				if err != nil {
					log.Default().Println(err)
				}

				refillRate := getRefillRate(newIntent.RiskScore)

				bucket.UpdateRefillRate(refillRate)
			}
		}

		bucketsMutex.Unlock()

		// Respond with 429 if rate limit exceeded
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
