package exectest

import "fmt"

//Big find the greatest among numbers
func Big() {
	data := []int{5, 7, 9, 20}

	fmt.Println(big(data))
}

func big(a []int) int {
	var s int

	for _, i := range a {
		if s < i {
			s = i

		}
	}
	return s
}
