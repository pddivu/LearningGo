package main

import "fmt"

// here are 2 call back functions
func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	xs := filter([]int{1, 2, 3, 4, 5}, func(a int) bool {
		return a > 1
	})
	fmt.Println(xs)
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(num []int, callback func(int) bool) []int {
	var xs []int
	for _, a := range num {
		if callback(a) {
			xs = append(xs, a)
		}
	}
	return xs
}
