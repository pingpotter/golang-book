package main

import "fmt"

func main() {
	fmt.Println("===============Zero Value====================")
	var number int
	var str string
	var boolean bool
	fmt.Printf("number: %v : '%v'\n", number, str)
	fmt.Printf("str: '%v'\n", str)
	fmt.Printf("line: %v\n", boolean)
}
