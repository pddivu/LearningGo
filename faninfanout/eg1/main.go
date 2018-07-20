package main

import (
	"fmt"
	"sync"
)

func main() {
	input := gen(2, 3)
	c1 := sq(input)
	c2 := sq(input)

	for q := range merge(c1, c2) {
		fmt.Println(q)
	}
}

func gen(num ...int) chan int {
	out := make(chan int)
	go func() {
		for n := range num {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(c chan int) chan int {

	out := make(chan int)

	go func() {
		for m := range c {
			out <- m * m
		}
		close(out)
	}()
	return out
}

func merge(a ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(a))

	for _, b := range a {
		go func(ch chan int) {
			for c := range ch {
				out <- c
			}
			wg.Done()
		}(b)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
