package concurrent

import (
	"fmt"
	"time"
)

func playTimer() {
	fmt.Println("===timer===")

	// call
	timer := time.NewTimer(time.Millisecond)
	<-timer.C
	fmt.Println("call")

	// stop
	timer = time.NewTimer(time.Millisecond)
	go func() {
		<-timer.C
		fmt.Println("call")
	}()
	ok := timer.Stop()
	if ok {
		fmt.Println("stop")
	}
}
