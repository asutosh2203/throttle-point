package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type ClientData struct {
	ReqTime  time.Time
	ReqCount int
}

var rateLimitMap sync.Map

const (
	limit  = 10
	window = 60 * time.Second
)

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.RemoteIP()

		now := time.Now()
		value, _ := rateLimitMap.LoadOrStore(ip, &ClientData{ReqTime: now, ReqCount: 1})
		clientData := value.(*ClientData)

		if time.Since(clientData.ReqTime) <= window {
			if clientData.ReqCount > limit {
				ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
				ctx.Abort()
				return
			}
			clientData.ReqCount++
		} else {
			clientData.ReqTime = now
			clientData.ReqCount = 1
		}

		ctx.Next()
	}
}
