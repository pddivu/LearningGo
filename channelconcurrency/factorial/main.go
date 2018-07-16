package main

import "fmt"

func main() {
	fact := factorial(2)
	for s := range fact {
		fmt.Println(s)
	}

}

func factorial(n int) chan int {
	c := make(chan int)

	f := 1
	go func() {
		for i := n; i > 0; i-- {
			f = f * i
		}
		c <- f
		close(c)
	}()
	return c
}
