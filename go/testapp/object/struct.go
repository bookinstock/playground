package object

import "fmt"

func play_struct() {
	fmt.Println("===struct===")

	u1 := NewUser("foo", 17)
	u2 := NewUser("bar", 18)
	fmt.Println("u1=", u1)
	fmt.Println("u2=", u2)

	r := rect{2, 2}
	fmt.Println("area=", r.area())
	fmt.Println("perim=", r.perim())
	fmt.Println("&area=", (&r).area())
	fmt.Println("&perim=", (&r).perim())
}

type rect struct {
	height, width int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2 * (r.height + r.width)
}
