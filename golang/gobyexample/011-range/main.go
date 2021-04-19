package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	// original code is "for _, num := range nums {"
	// Key, Value
	for test, num := range nums {
		// returns key
		fmt.Println("key:", test)
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		// returns Unicode code points
		// 0, 103
		// 1, 111
		fmt.Println(i, c)
	}
}
