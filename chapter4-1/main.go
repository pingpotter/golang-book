package main

import (
	"time"
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

	fmt.Printf("Value : %v Type: %T\n", 3.2111, 3.2111)
	fmt.Printf("Value : %v Type: %T\n", true, true)
	fmt.Printf("Value : %v Type: %T\n", time.Now(), time.Now())
}
