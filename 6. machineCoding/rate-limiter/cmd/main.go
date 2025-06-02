package main

import (
	"fmt"
	"rate-limiter/limiter"
	"time"
)

func main() {
	// Allow 5 requests per second
	rl := limiter.NewRateLimiter(5, 5)

	for i := 1; i <= 10; i++ {
		allowed := rl.Allow()
		fmt.Printf("Request %d allowed: %v\n", i, allowed)
		time.Sleep(100 * time.Millisecond) // simulate delay
	}
}