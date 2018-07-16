package main

import "fmt"

func main() {
	c := que(2, 3, 5, 9)
	sq := square(c)
	/*fmt.Println(<-sq)
	fmt.Println(<-sq)*/

	for m := range sq {
		fmt.Println(m)
	}
}

func que(num ...int) chan int {

	out := make(chan int)

	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(ch chan int) chan int {
	out := make(chan int)
	go func() {
		for p := range ch {
			out <- p * p
		}
		close(out)
	}()
	return out
}
