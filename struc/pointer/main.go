package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := &person{"Divya", "Divakaran", 34}
	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1.first)
}
