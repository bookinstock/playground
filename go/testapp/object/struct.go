package object

import "fmt"

func play_struct() {
	fmt.Println("===struct===")

	u1 := NewUser("foo", 17)
	u2 := NewUser("bar", 18)
	fmt.Println("u1=", u1)
	fmt.Println("u2=", u2)

	r := Rect{2, 2}
	fmt.Println("area=", r.area())
	fmt.Println("perim=", r.perim())
	fmt.Println("&area=", (&r).area())
	fmt.Println("&perim=", (&r).perim())
}

type Rect struct {
	height, width int
}

func (r Rect) area() int {
	return r.height * r.width
}

func (r *Rect) perim() int {
	return 2 * (r.height + r.width)
}
