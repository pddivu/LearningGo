package main

//copies a slice
import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4}

	check := a
	// copy set the value of a and check as b
	copy(a, b)
	//copy(b, a)
	//a = b  -> changes the value of a but not value of check
	fmt.Println(a, b, check)
}
