package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 56, 6}
	fmt.Println(slice)
	newSlice := make([]int, 2)
	fmt.Println(newSlice)
	copy(slice, newSlice)
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Slice: %v\n", newSlice)

}
