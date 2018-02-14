package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {

		fmt.Println(f(i))
	}
}

func f(i int) (int, string) {

	y := ""

	if i%15 == 0 {
		y = "FizzBuzz"
	} else if i%3 == 0 {
		y = "Fizz"
	} else if i%5 == 0 {
		y = "Buzz"
	}

	return i, y
}
