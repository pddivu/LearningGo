package main

import "fmt"

/*  Binary number of entered digit and count of 1 in binary*/

func main() {
	var num, k, l int

	fmt.Println("Enter a digit")
	fmt.Scan(&k)
	l = k

	for {
		if k%2 == 1 {
			num = num + 1
		}
		k = k / 2
		if k == 0 {
			break
		}
	}
	fmt.Printf("Number - %d, Binary -%b, Count of 1 s- %d\n", l, l, num)
}
