package main

import "fmt"

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
}
