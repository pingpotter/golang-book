package main

import (
	"fmt"
)

func main() {

	x := map[int]int{
		1: 1,
		2: 2,
		//3: 0,
	}

	fmt.Println(x[3])

	if x[3] != 0 {
		fmt.Println(x[3])

	}

	///ok
	if value, ok := x[3]; ok {
		fmt.Println(value)
	}
}
