package concurrent

import (
	"fmt"
	"time"
)

const introChannelSync = `
use channels to synchronize execution across goroutines (single)
using a blocking receive to wait for a goroutine to finish.
waiting for multiple goroutines to finish (WaitGroup)
`

func playChannelSync() {
	fmt.Println("===playChannelSync===")
	fmt.Println(introChannelSync)

	queue := make(chan int)
	go func() {
		fmt.Println("work start")
		time.Sleep(time.Millisecond)
		queue <- 1
		fmt.Println("work finish")
	}()

	fmt.Println("wait")
	<-queue
}
