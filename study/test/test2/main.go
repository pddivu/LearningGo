package main

import "fmt"

func main() {

	a := 1234567

	var b, c int

	for k := 0; a != 0; k++ {
		b = a % 10
		a = a / 10
		c = c + b

		c = c * 10

		fmt.Println(b, c)
	}
	fmt.Println(c / 10)

}
