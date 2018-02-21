package main

import (
	"fmt"
)

type Circle struct {
	x float64
	y float64
	z float64
}

func main() {
	var c Circle
	fmt.Printf("c type: %T\n", c)
	fmt.Println(c.x, c.y, c.z)

	c1 := new(Circle)
	fmt.Printf("c1 type: %T\n", c1)
	fmt.Println(c1.x, c1.y, c1.z)

	c2 := Circle{x: 0, y: 1, z: 5}
	fmt.Printf("c2 type: %T\n", c2)
	fmt.Println(c2.x, c2.y, c2.z)

	c3 := NewCircle(1, 2, 3)
	fmt.Printf("c3 type: %T\n", c3)
	fmt.Println(c3.x, c3.y, c3.z)
}

func NewCircle(i, j, k float64) *Circle {
	return &Circle{i, j, k}
}
