package main

import (
	"fmt"
)

func main() {

	y := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, number := range y {
		fmt.Println(key, number)
	}

}
