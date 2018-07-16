package main

import "fmt"

func main() {
	for p := range fact(numlist()) {
		fmt.Println(p)
	}
}

func numlist() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func fact(f chan int) chan int {
	out := make(chan int)

	go func() {
		for q := range f {
			l := 1

			for r := q; r > 0; r-- {
				l = l * r

			}
			out <- l
		}
		close(out)
	}()
	return out
}
