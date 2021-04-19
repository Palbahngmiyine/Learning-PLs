package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	// same as test := m["k2"]
	test, _ := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("test:", test)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
