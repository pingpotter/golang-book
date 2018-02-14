package main

import (
	"fmt"
)

func main() {

	number := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(number); i++ {
		fmt.Println(i, number[i])

	}
	fmt.Println("With Range")

	for i, number := range number {
		fmt.Println(i, number)
	}
}
