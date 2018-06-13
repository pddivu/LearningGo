package main

import "fmt"

var a string

func main() {
	fmt.Println("Enter a word")
	fmt.Scanln(&a)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
