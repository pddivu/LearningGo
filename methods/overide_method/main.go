package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type personality struct {
	person
	wish string
}

func (p person) greeting() {
	fmt.Println(p)
	fmt.Println("I am regular person")
}
func (q personality) greeting() {
	fmt.Println(q)
	fmt.Println("I am a Batman")
}

func main() {
	p1 := person{"Divya", "Divakaran", 34}
	p2 := personality{person{"Eshaan", "Bhagavad", 7}, "Good boy"}

	p1.greeting()
	p2.greeting()
	p2.person.greeting()

}
