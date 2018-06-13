package main

import "fmt"

var num int
var st string

func main() {

	var va int
	va = 1
	fmt.Println("var -", va)
	fmt.Println(&va)
	num = 5

	//ad1 is a pointer/address decalration
	var ad1 *int = &num
	*ad1 = 6
	fmt.Println(ad1)
	va = comp()
	fmt.Println("var -", va)
	fmt.Println(&va)
}
func comp() int {
	var va int
	va = num + 1
	fmt.Println("var -", va)
	fmt.Println(&va)
	return (va)
	//everything is pass by value in Go
}
