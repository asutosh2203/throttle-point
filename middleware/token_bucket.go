package middleware

import (
	"log"
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity       float64
	Tokens         float64
	RefillRate     float64 // tokens per second
	RequestCount   int64
	LastRefillTime time.Time
	Mutex          sync.Mutex
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		Capacity:       capacity,
		Tokens:         capacity,
		RefillRate:     refillRate,
		LastRefillTime: time.Now(),
		RequestCount:   0,
	}
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.Mutex.Lock()
	defer tb.Mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.LastRefillTime).Seconds()

	tb.Tokens = math.Min(tb.Capacity, tb.Tokens+elapsed*tb.RefillRate)
	tb.LastRefillTime = now

	if tb.Tokens >= 1 {
		tb.Tokens -= 1
		log.Println("Tokens Remaining: ", tb.Tokens)
		return true
	}

	return false
}

func (tb *TokenBucket) UpdateRefillRate(newRefillRate float64) {
	tb.Mutex.Lock()
	defer tb.Mutex.Unlock()
	tb.RefillRate = newRefillRate
}
