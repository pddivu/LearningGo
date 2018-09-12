package main

import "fmt"

type currency int

/*USD is currency*/
const (
	USD currency = iota
	GBP
	EUR
	RMB
)

func main() {
	var x [100]int
	/* array with total capacity of 100 with length 50

	a slice with capacity of 100 with length 50
	*/
	array := new([100]int)[0:50]
	slice := make([]int, 50, 100)
	fmt.Println(slice, array)
	for i := 0; i < 100; i++ {
		x[i] = i
	}
	fmt.Println(x)
	fmt.Println(x[5])
	fmt.Println(x[2:6])
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))

	symbols := [...]string{USD: "$", GBP: "@", EUR: "#", RMB: "%"}
	fmt.Println(symbols)
	fmt.Println(USD, symbols[USD])

	// Creates an array of 10 ie, 0-9 and 9th value will be -1
	num := [...]int{9: -1}
	fmt.Println(num)

	months := [...]string{1: "January", 2: "February", 3: "March" /*  */, 12: "December"}
	summer := months[2:3]
	fmt.Println(months[0:])
	fmt.Printf("%T\n", months)
	fmt.Println(summer)

}
