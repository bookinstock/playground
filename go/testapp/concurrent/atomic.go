package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const introAtomic = `
managing state
using the sync/atomic package for atomic counters accessed by multiple goroutines.`

func playAtomic() {
	fmt.Println("===atomic===")
	fmt.Println(introAtomic)

	var wg sync.WaitGroup

	var count uint64

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			// count += 1

			atomic.AddUint64(&count, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("count=", count)
}
