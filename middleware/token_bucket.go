package middleware

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity       float64
	Tokens         float64
	RefillRate     float64 // tokens per second
	LastRefillTime time.Time
	Mutex          sync.Mutex
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		Capacity:       capacity,
		Tokens:         capacity,
		RefillRate:     refillRate,
		LastRefillTime: time.Now(),
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
		return true
	}

	return false
}
