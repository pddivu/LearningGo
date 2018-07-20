package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type skill struct {
	person
	skilltocode bool
}

func main() {
	p1 := skill{
		person: person{
			First: "Divya",
			Last:  "Divakaran",
			Age:   34,
		},
		skilltocode: true,
	}

	p2 := skill{person{"Eshaan", "Bhagavad", 6}, false}

	fmt.Println(p1)
	fmt.Println(p2)
}
