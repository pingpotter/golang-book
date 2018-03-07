package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(16)

	for i := 1; i <= 16; i++ {
		go increment(i)
	}
	wg.Wait()
	fmt.Println("Finished ", counter)
}

func increment(n int) {

	defer wg.Done()
	for x := 0; x < 2; x++ {
		value := counter
		//runtime.Gosched()
		value++
		counter = value
	}
}
