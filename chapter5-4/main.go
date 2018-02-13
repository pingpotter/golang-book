package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	random := rand.Intn(100)
	fmt.Println("Random : ", random)

	for i := 0; i < 5; i++ {

		var input int
		fmt.Scanf("%d\n", &input)

		if input > random {
			fmt.Println(input, "มากไป")
		} else if input < random {
			fmt.Println(input, "น้อยไป")
		} else {
			fmt.Println(input, "เจอแล้ว")
			return
		}

	}
}
