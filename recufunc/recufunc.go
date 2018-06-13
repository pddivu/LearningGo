package main

import "fmt"

var f int

func main() {

	var a int
	fmt.Println("Enter a Number")
	fmt.Scan(&a)
	f = 1
	f = Fact(a)
	// Call Fact Function with out an receiving variable
	fmt.Println(Fact(a))
	fmt.Printf("The Factorial of %d is %d \n", a, f)
}

// "Fact- Factorial of entered number"
func Fact(n int) int {
	//var b int
	if n == 0 {
		return 1
		// code will be ended here
	}
	return n * Fact(n-1)

}
