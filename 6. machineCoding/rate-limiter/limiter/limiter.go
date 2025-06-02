package limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate       int           // tokens per second
	capacity   int           // max tokens
	tokens     float64       // current tokens
	lastRefill time.Time     // last time tokens were added
	mu         sync.Mutex
}

func NewRateLimiter(rate int, capacity int) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		capacity:   capacity,
		tokens:     float64(capacity),
		lastRefill: time.Now(),
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastRefill).Seconds()
	// Refill tokens based on time passed
	refill := elapsed * float64(r.rate)
	if refill > 0 {
		r.tokens = min(float64(r.capacity), r.tokens+refill)
		r.lastRefill = now
	}

	if r.tokens >= 1 {
		r.tokens--
		return true
	}

	return false
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}