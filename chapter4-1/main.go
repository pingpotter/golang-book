package main

import (
	"fmt"
)

func main() {

	//Long
	var x string = "Hello World 1"
	fmt.Println(x)

	var y string
	y = "Hello World 2"
	fmt.Println(y)

	//Short
	z := "Hello World 3"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)
}
