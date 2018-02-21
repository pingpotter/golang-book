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

type Rectangle struct {
	w float64
	h float64
}

type measure interface {
	area() float64
}

func main() {

	bigC := &Circle{0, 0, 5}
	printArea(bigC)

	r := Rectangle{3, 4}
	printArea(r)
}

func printArea(m measure) {
	fmt.Println(m.area())
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func (c *Circle) area() float64 {
	return math.Pi * c.z * c.z
}
