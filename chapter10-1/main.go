package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	z float64
}

func main() {
	littleC := Circle{0, 0, 5}
	fmt.Println("LittleC", littleC.area())
	littleC.changeRedius(10)
	fmt.Println("LittleC", littleC.area())

	bigC := &Circle{0, 0, 5}
	fmt.Println("bigC", bigC.area())
	bigC.changeRedius(10)
	fmt.Println("bigC", bigC.area())
}

func (c Circle) area() float64 {
	return math.Pi * c.z * c.z
}

func (c *Circle) changeRedius(r float64) {
	c.z = r
}
