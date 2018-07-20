package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

//First is variable name in bith structure
type personality struct {
	person
	First string
	good  bool
}

func main() {
	p1 := personality{person{"Divya", "Divakaran", 34}, "Good", true}
	fmt.Println(p1)
	//First is variable name in bith structure
	fmt.Println(p1.First, p1.person.First)
}
