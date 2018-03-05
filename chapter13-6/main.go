package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
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
		atomic.AddInt64(&counter, 1)
	}
}
