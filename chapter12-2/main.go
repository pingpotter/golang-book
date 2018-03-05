package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		func(a, b int) int {
			return a + b
		}(2, 3))

	fmt.Println(add(2, 53))
}

func add(a, b int) int {
	return a + b

}
