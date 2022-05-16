package concurrent

import (
	"fmt"
	"time"
)

const introTicker = "tickers are for when you want to do something repeatedly at regular intervals"

func playTicker() {
	fmt.Println("===ticker===")
	fmt.Println(introTicker)

	ticker := time.NewTicker(time.Millisecond * 100)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("call")
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()

	time.Sleep(time.Millisecond * 500)
	ticker.Stop()
	time.Sleep(time.Millisecond * 500)
	done <- true
}
