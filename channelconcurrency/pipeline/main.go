package main

import "fmt"

func main() {

	//sq := square(square(que(2, 3, 4, 5, 6)))
	for p := range square(square(que(2, 3, 4, 5, 6))) {
		fmt.Println(p)
	}
}

func que(num ...int) chan int {
	out := make(chan int)

	go func() {
		for _, r := range num {
			out <- r
		}
		close(out)
	}()
	return out
}

func square(m chan int) chan int {

	out := make(chan int)

	go func() {
		for q := range m {
			out <- q * q
		}

		close(out)
	}()
	return out
}
