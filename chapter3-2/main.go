package main

import "fmt"

func main() {
	fmt.Println("===============String====================")
	backticks := `hello world,
	today's good day.`
	fmt.Println(backticks)

	doubleQuote := "hello world,\ntoday's bad day."
	fmt.Println(doubleQuote)
}
