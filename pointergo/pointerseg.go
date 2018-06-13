package main

import "fmt"

func main() {

	x := 5
	zero(x)
	fmt.Println(x)
	fmt.Printf("%p\n", &x)

	zero1(&x)
	fmt.Println(x)
	fmt.Printf("%p\n", &x)

}

// x gets the value zero but not returning 0
func zero(x int) {
	x = 0
	fmt.Printf("%p\n", &x)

}

//address is passed and the value is set
func zero1(x *int) {
	*x++
	fmt.Printf("%p\n", &x)

}
