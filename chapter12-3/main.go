package main

import (
	"fmt"
)

func main() {
	addFunc := func(a int) func(b int) int {
		fmt.Println(a)
		return func(b int) int {
			fmt.Println(b)
			return a + b

		}
	}

	addFuncNew := addFunc(2)
	fmt.Println(addFuncNew(3))
	addFuncNew2 := addFunc(7)
	fmt.Println(addFuncNew2(3))
}
