package main

import "fmt"

func main() {

	var a = []int{1, 2, 3, 4}

	swapnum(a)
	fmt.Println(a)
}

func swapnum(n []int) {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {

		n[i], n[j] = n[j], n[i]
	}
}
