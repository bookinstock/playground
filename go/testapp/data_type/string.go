package data_type

import (
	"fmt"
	"unicode/utf8"
)

func playString() {

	fmt.Println("===string===")
	fmt.Println("a+b=", "a"+"b")

	const s = "你好"
	fmt.Println("len s=", len(s))
	fmt.Println("len rune s=", utf8.RuneCountInString(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	for i, e := range s {
		fmt.Printf("%#U starts at %d\n", e, i)
	}
	fmt.Println()

}
