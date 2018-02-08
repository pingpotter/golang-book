package main

import (
	"fmt"
)


func main() {

	fmt.Print("Enter a number : ")

	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * (5.00/9.00)
	fmt.Printf("%.2f",output)

}
