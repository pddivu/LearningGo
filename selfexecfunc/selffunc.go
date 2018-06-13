package main

import "fmt"

//SElf executing function

func main() {
	func() {
		fmt.Println("I am Happy!")
	}() // this paranthesis makes it execute
}
