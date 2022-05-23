package concurrent

import "fmt"

const introAtomic = `
managing state
using the sync/atomic package for atomic counters accessed by multiple goroutines.`

func playAtomic() {
	fmt.Println("===atomic===")
	fmt.Println(introAtomic)

}
