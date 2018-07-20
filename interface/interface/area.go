package main

import (
	"fmt"
	"math"
)

type square struct {
	side float32
}

type circle struct {
	rad float32
}

type shape interface {
	area() float32
}

func (z square) area() float32 {
	return z.side * z.side
}

func (c circle) area() float32 {
	return math.Pi * c.rad * c.rad
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())

}
func totalarea(shapes ...shape) float32 {
	var area float32
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func main() {
	s := square{10}
	c := circle{10}
	info(s)
	info(c)
	fmt.Println("Totalarea:", totalarea(s, c))
}
