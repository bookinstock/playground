package data_type

import "fmt"

func play_map() {
	m := make(map[string]int)
	fmt.Println("m=", m)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println("m=", m)

	fmt.Println("len=", len(m))

	delete(m, "b")
	fmt.Println("m=", m)

	_, prs := m["b"]
	fmt.Println("prs=", prs)

	m = map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
	}
	fmt.Println()
}
