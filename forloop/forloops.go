package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i, "-", j)
		}
	}

	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 10 {
			break
		}

	}

	for a := 1; a < 15; a++ {
		if a%2 == 0 {
			fmt.Printf("%d is even number\n", a)
		} else {
			fmt.Printf("%d is odd number\n", a)

		}
	}
}
