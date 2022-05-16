package concurrent

import (
	"fmt"
	"time"
)

func playWorkerPool() {
	fmt.Println("+++worker pool+++")

	jobCount := 10
	workerCount := 3

	jobs := make(chan int, jobCount)
	results := make(chan int, jobCount)

	worker := func(workerId int) {
		fmt.Println("Worker start", workerId)
		for e := range jobs {
			fmt.Println("Worker processing", workerId, e)
			time.Sleep(time.Microsecond)
			results <- e * 2
		}
		fmt.Println("Worker stop", workerId)
	}

	for i := 0; i < workerCount; i++ {
		go worker(i)
	}

	for i := 0; i < jobCount; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < jobCount; i++ {
		fmt.Println("results=", <-results)
	}
}
