package main

import "fmt"

func main() {
	if true {
		fmt.Println("1 - This is true")

	}
	if false {
		fmt.Println("1- This is false")

	}
	/////////////
	if !true {
		fmt.Println("2 - This is false")

	}
	if !false {
		fmt.Println("2 - This is true")

	}
	////////////////

	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	/////////////

	number := 13
	if number == 15 {
		fmt.Println("Fifteen")
	} else if number == 14 {
		fmt.Println("Fourteen")

	} else if number == 13 {
		fmt.Println("Thirteen")
	} else {
		fmt.Println("No number")
	}
	////////
}
