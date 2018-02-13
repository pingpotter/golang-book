package main

import (
	"fmt"
	"time"
)

var j string = "Hello World 5"

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

	const a string = "Hello World"
	//a = "Other string"

	var (
		f = 5
		g = 10
		h = 15
	)

	fmt.Printf("%v %v %v\n", f, g, h)

	v1, v2 := "First", "Second"
	fmt.Printf("%v %v\n", v1, v2)

	v1, v2 = v2, v1
	fmt.Printf("%v %v\n", v1, v2)
}

func f() {

	fmt.Println(j)
}
