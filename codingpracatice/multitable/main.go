package main

import (
	"fmt"
)

//Print the multiplication table of N upto 10.
func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Enter a number greater than 0")
	} else {
		for i := 1; i < 11; i++ {
			fmt.Println(n * i)
		}
	}
}
