package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i, f(i))
	}
}

func f(i int) string {
	/*FizzBuzz := []string{strconv.Itoa(i), "Fizz", "Buzz"}
	start := 0
	end := 1
	if i%3 == 0 {
		start = 1
		end = 2
	}
	if i%5 == 0 {
		start = 2
		end = 3
	}*/
	FizzBuzz := []string{"F", "i", "z", "z", "B", "u", "z", "z"}
	//start := i*i%3*4
	//end := 8-((-(10**4))%5)
	start := ((i * i) % 3) * 4
	x := int(math.Pow(float64(i), 4.00))
	end := 8 - (-(x) % 5)

	if end > 8 {
		end = 0
	}
	fmt.Println(start)
	fmt.Println(end)

	return strings.Join(FizzBuzz[start:end], ",")
}
