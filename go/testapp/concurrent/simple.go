package concurrent

import (
	"fmt"
	"time"
)

const introSimple = "A goroutine is a lightweight thread of execution."

func playSimple() {
	fmt.Println("===simple===")

	go foo()

	go func() {
		fmt.Println("bar")
	}()

	fmt.Println(introSimple)

	time.Sleep(time.Millisecond)

}

func foo() {
	fmt.Println("foo")
}
