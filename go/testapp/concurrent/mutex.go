package concurrent

import (
	"fmt"
	"sync"
)

func playMutex() {
	fmt.Println("===mutex===")

	playMutexCount()

	playMutexMap()
}

func playMutexCount() {
	type Counter struct {
		val  uint
		lock sync.Mutex
	}

	var wg sync.WaitGroup

	var counter Counter

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			counter.lock.Lock()
			defer counter.lock.Unlock()
			defer wg.Done()
			counter.val += 1
		}()
	}

	wg.Wait()

	fmt.Println("counter=", counter.val)
}

func playMutexMap() {

	type Counter struct {
		val  map[string]int
		lock sync.Mutex
	}

	var wg sync.WaitGroup

	counter := Counter{val: map[string]int{}}

	letters := []string{"a", "b", "c"}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			counter.lock.Lock()
			defer counter.lock.Unlock()
			defer wg.Done()

			letter := letters[n%3]

			counter.val[letter] += 1
		}(i)
	}

	wg.Wait()

	fmt.Println("counter=", counter.val)
}
