package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {

		switch i {
		case 0:
			fmt.Println(i, "zero")
		case 1:
			fmt.Println(i, "one")
		case 2:
			fmt.Println(i, "two")
		case 3:
			fmt.Println(i, "three")
		case 4:
			fmt.Println(i, "Four")
		default:
			fmt.Println(i)
		}

	}
}
