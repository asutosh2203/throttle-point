package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ClientData struct {
	ReqTime  time.Time
	ReqCount int
}

// rate limit map
var rateLimitMap = make(map[string]*ClientData)

func main() {
	fmt.Println("Welcome to ThrottlePoint")
	r := gin.Default()

	r.Any("/*proxyPath", func(ctx *gin.Context) {

		// rate limiting logic - allowing 10 reqs per minute
		ip := ctx.RemoteIP()
		clientData := rateLimitMap[ip]

		if clientData != nil {
			if time.Since(clientData.ReqTime) <= 60*time.Second {
				if clientData.ReqCount >= 10 {
					ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "Cool your jets mate"})
					return
				} else {
					clientData.ReqCount++
				}
			} else {
				clientData.ReqTime = time.Now()
				clientData.ReqCount = 1
			}
		} else {
			newClient := &ClientData{ReqTime: time.Now(), ReqCount: 1}
			rateLimitMap[ip] = newClient
		}

		fmt.Println()
		fmt.Println(rateLimitMap[ip])
		fmt.Println()

		// construct target URL
		targetURL := "http://localhost:3000" + ctx.Request.URL.Path
		if ctx.Request.URL.RawQuery != "" {
			targetURL += "?" + ctx.Request.URL.RawQuery
		}

		// construct new http request towards target
		req, err := http.NewRequest(ctx.Request.Method, targetURL, ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
			return
		}

		// copy headers into the new request
		for key, values := range ctx.Request.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		// send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": "failed to reach backend"})
			return
		}

		defer resp.Body.Close()

		// Set response headers
		for key, values := range resp.Header {
			for _, value := range values {
				ctx.Writer.Header().Add(key, value)
			}
		}

		ctx.Status(resp.StatusCode)
		io.Copy(ctx.Writer, resp.Body)

	})

	// Start the server
	r.Run("0.0.0.0:8080")
}
