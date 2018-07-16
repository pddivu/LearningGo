package main

import "fmt"

func main() {

	var x []string

	y := []string{"a", "b", "c", "d"}

	x = []string{"a"}
	fmt.Println(y, x)
	fmt.Println(y[2:4])
	//Need to learn on myString
	fmt.Println("myString"[0])
	slice := make([]int, 0, 5)
	fmt.Println("--------------------------------")
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println("--------------------------------")
	for i := 0; i < 40; i++ {
		slice = append(slice, i)
		fmt.Println(slice[i], " len-", len(slice), " Cap-", cap(slice))
	}

	//Appending Slice with slice elements again , like variadic expression

	slice = append(slice, slice...)
	fmt.Println(slice)
}
