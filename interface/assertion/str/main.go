package main

import "fmt"

func main() {
	var name interface{} = "Divya"

	// string assertion
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Println("Name is not string")
	}
	var value interface{} = 7
	fmt.Printf("%T\n", value)

	// int assertion
	fmt.Println(value.(int) + 6)

	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))
}
