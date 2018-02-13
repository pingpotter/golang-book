package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 1; ; i++ {

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		random := r1.Intn(100)
		fmt.Println(i, "Random : ", random)

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

		if i > 5 {
			return
		}
	}
}
