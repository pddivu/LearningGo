package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(3)

	var count int

	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println("Total count - ", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {

		go func(i int) {
			for k := 0; k < n; k++ {
				c <- fmt.Sprintf("Process "+strconv.Itoa(i)+"printing: ", k)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	return c
}
