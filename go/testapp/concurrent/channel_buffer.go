package concurrent

import "fmt"

const introChannelBuffer = `
By default channels are unbuffered,
meaning that they will only accept sends (chan <-)
if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.
`

func playChannelBuffer() {
	fmt.Println("===channel buffer===")
	fmt.Println(introChannelBuffer)

}
