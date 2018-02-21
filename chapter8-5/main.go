package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	double(slice)
	fmt.Printf("Original addr %p\n", slice)
	fmt.Printf("Original %v\n", slice)
}

func double(number []int) {
	fmt.Printf("double addr %p\n", number)
	fmt.Printf("double value %v\n", number)
	for i := range number {
		number[i] *= 2
	}
	fmt.Println(number)
}
