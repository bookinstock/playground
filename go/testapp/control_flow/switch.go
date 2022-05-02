package control_flow

import (
	"fmt"
	"time"
)

func play_switch() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("beforenoon")
	default:
		fmt.Println("afternoon")
	}

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool", t)
		case int:
			fmt.Println("int", t)
		default:
			fmt.Println("unknown", t)
		}
	}

	checkType(true)
	checkType(100)
	checkType("foo")
}
