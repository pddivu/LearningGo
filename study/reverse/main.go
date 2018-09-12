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

	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	copy(a, b)
	fmt.Println(a, b, check)

	a1 := []int{1, 2}
	b1 := []int{3, 4}
	check1 := a1
	a1 = b1
	fmt.Println(a1, b1, check1)

	a2 := map[string]bool{"A": true, "B": true}
	b2 := make(map[string]bool)
	for key, value := range a2 {
		b2[key] = value
	}

	a3 := map[string]bool{"A": true, "B": true}
	b3 := map[string]bool{"C": true, "D": true}
	check3 := a3
	a3 = b3
	fmt.Println(a3, b3, check3)

}
