package main

import (
	"fmt"
)

func main() {
	m := map[int]int{1:1, 2:2, 3:3}
	double(m)
	fmt.Printf("Original addr %p\n", m)
	fmt.Printf("Original %v\n", m)
}

func double(number map[int]int) {
	fmt.Printf("double addr %p\n", number)
	fmt.Printf("double value %v\n", number)
	for i := range number {
		number[i] *= 2
	}
	fmt.Println(number)
}
