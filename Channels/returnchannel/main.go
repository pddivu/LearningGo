package main

import "fmt"

func main() {

	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)

	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for s := range c {
			sum += s
		}
		out <- sum
		close(out)
	}()
	return out
}
