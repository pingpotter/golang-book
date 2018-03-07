package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < 5; i++ {
		go func(n int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(n, ":", i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished")
}

func f(n int) {

	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
