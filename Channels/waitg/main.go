package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wt sync.WaitGroup
	wt.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wt.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wt.Done()
	}()
	go func() {
		wt.Wait()
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
}
