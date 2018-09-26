package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := [3]int{1, 2, 3}

	fmt.Println(x, y)
	double(x)      // 2 4 6
	fmt.Println(x) // 2 4 6
	double1(y)     // 2 4 6
	fmt.Println(y) // 1 2  3
}

func double(s []int) {

	for i := 0; i < 3; i++ {
		s[i] = s[i] * 2

	}
	fmt.Println(s)
}
func double1(s [3]int) {

	for i := 0; i < 3; i++ {
		s[i] = s[i] * 2

	}
	fmt.Println(s)
}
