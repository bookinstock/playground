package concurrent

import "fmt"

func Run() {
	fmt.Println("###concurrent###")

	playSimple()

	playChannel()

	playChannelBuffer()

	playChannelSync()

	playChannelDirection()

	playChannelSelect()

	playTimeout()

	playChannelNonBlocking()

	playChannelClose()

	playRange()
}
