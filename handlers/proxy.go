package handlers

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(ctx *gin.Context) {

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
	client := &http.Client{Timeout: 10 * time.Second}

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

}
