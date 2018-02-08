package main

import "fmt"

func main() {
	fmt.Println("===============Floating Point====================")
	third := 1.0 / 3.0
	fmt.Printf("third = %v\n", third)
	fmt.Printf("third + third + third = %v\n", third+third+third)

	fmt.Println("==============Comparing floating point============")

	fmt.Println("0.1 + 0.2 == 0.3 is ", 0.1+0.2 == 0.3)
	var num float32
	num = 0.1
	num += 0.2
	fmt.Println("num == 0.3 is ", num == 0.3)
	fmt.Println("num is ", num)
	//https://stackoverflow.com/questions/22337418/golang-floating-point-precision-float32-vs-float64
}
