package concurrent

import (
	"fmt"
	"sync"
	"time"
)

const introWaitGroup = "To wait for multiple goroutines to finish"

func playWaitGroup() {
	fmt.Println("===wait group===")
	fmt.Println(introWaitGroup)

	worker := func(workerId int) {
		fmt.Println("start worker", workerId)
		time.Sleep(time.Millisecond)
		fmt.Println("stop worker", workerId)
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			worker(workerId)
		}(i)
	}

	wg.Wait()

}
