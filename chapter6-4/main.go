package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2, 3))

	xs := []int{4, 5, 6}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0

	for _, v := range args {
		total += v
	}

	return total
}
