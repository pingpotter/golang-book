package main

import (
	"fmt"
)

func main() {
	amount := 5
	double(&amount)
	fmt.Printf("Original %v\n", amount)
}

func double(number *int) {
	*number *= 2
	fmt.Println(*number)
}
