package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 10)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[5] = 34

	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	fmt.Println("Slice with length and capacity")
	fmt.Printf("Slice length %v ,capacity %v, %v\n", len(slice), cap(slice), slice)

	//append
	for i := 4; i < 15; i++ {
		slice = append(slice, i)

	}

	fmt.Printf("Slice length %v ,capacity %v, %v\n", len(slice), cap(slice), slice)

}
