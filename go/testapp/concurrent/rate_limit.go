package concurrent

import (
	"fmt"
	"time"
)

const introRateLimit = `
Rate limiting is an important mechanism for
controlling resource utilization and maintaining quality of service.
Go elegantly supports rate limiting with goroutines, channels, and tickers.`

func playRateLimit() {
	fmt.Println("===rate limit===")

	fmt.Println(introRateLimit)

	rateLimit()

	fmt.Println("---")

	rateLimitBurst()
}

func rateLimit() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.NewTicker(200 * time.Millisecond)

	for req := range requests {
		<-limiter.C
		fmt.Println("req=", req, time.Now())
	}
}

func rateLimitBurst() {
	limiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		limiter <- time.Now()
	}

	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		for t := range ticker.C {
			limiter <- t
		}
	}()

	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-limiter
		fmt.Println("req=", req, time.Now())
	}

}
