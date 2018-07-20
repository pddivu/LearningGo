package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullname() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"Divya", "Divakaran", 34}
	p2 := person{"Eshaan", "Bhgavad", 7}

	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
