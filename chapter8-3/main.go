package main

import (
	"fmt"
)

func main() {
	array := [3]int{1, 2, 3}
	double(array)
	fmt.Printf("Original addr %p\n", &array)
	fmt.Printf("Original %v\n", array)
}

func double(number [3]int) {
	fmt.Printf("double addr %p\n", &number)
	for i := range number {
		number[i] *= 2
	}
	fmt.Println(number)
}
