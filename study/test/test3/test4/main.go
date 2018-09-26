package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p person) change() {
	p.Name = "Divya"
	p.Age = 32
	fmt.Println(p.Name, p.Age)
}

func main() {
	var p1 person
	fmt.Println(p1.Name, p1.Age) // Space and 0
	p1.change()                  // Divya 32
	fmt.Println(p1.Name, p1.Age) // Space and 0

}
