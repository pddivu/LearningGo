package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type que int

func main() {

	P1 := person{"Divya", "Divakaran", 34}
	P2 := person{"Eshaan", "Bhagavad", 6}
	fmt.Println(P2)
	fmt.Println(P1.first)

	var num que
	num = 12
	fmt.Printf("%T  %v\n", num, num)

}
