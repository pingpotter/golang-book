package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 1

	for i := 1; i <= 100; i++ {

		fmt.Println(i, f(i))
	}
}

func f(num int) string {

	ln := [3]int{15, 3, 5}
	str := [3]string{"FizzBuzz", "Fizz", "Buzz"}

	for i := 0; i < len(ln); i++ {
		if num%ln[i] == 0 {
			return str[i]
		}
	}

	return strconv.Itoa(num)
}
