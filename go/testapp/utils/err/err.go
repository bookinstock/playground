package err

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	fmt.Println("===os===")

	interrupt := make(chan os.Signal, 1)

	go goexit()

	// TODO not work
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	sig := <-interrupt

	fmt.Println("sig=", sig)
}

func goexit() {
	time.Sleep(time.Second)
	os.Exit(1)
}

func gopanic() {
	time.Sleep(time.Second)
	panic("foo")
}
