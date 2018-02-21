package main

import (
	"fmt"
)

type Zipcode string

func main() {
	zipcode := Zipcode("11111")
	if zipcode.valid() {
		fmt.Println(zipcode)
	}
}

func (z Zipcode) valid() bool {
	return true
}
