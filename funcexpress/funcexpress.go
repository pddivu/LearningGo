package main

import "fmt"

//func expreseeion is passing function to a variable
func main() {

	greeting := func() {
		fmt.Println("Hi Divya")
	}

	greeting()
	fmt.Printf("%T\n", greeting)

	name := who()
	fmt.Println(name())
	fmt.Printf("%T\n", name)
}

func who() func() string {
	return func() string {
		fmt.Println("I am Divya")
		return "Hi Divya"
	}
}
