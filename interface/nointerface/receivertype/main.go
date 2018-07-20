package main

import (
	"fmt"
	"math"
)

type circle struct {
	radious float32
}

type square struct {
	side float32
}

type shape interface {
	area() float32
}

func (c *circle) area() float32 {
	return math.Pi * c.radious
}

//func (s *square) area() float32 {
func (s square) area() float32 {
	return s.side * s.side
}

func info(p shape) {
	fmt.Println("area:", p.area())
}

func main() {
	a := circle{10}
	b := square{20}

	//pointer receiver can accept a pointer or a value
	// vaalue receiver can accept only value
	info(&a)
	info(&b)
}
