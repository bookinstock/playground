package concurrent

import "fmt"

func Run() {
	fmt.Println("###concurrent###")

	playSimple()

	playChannel()

	playChannelBuffer()
}
