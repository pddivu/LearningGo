package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")
	time.Sleep(4000 * time.Millisecond) // if we did not put this it will not
	// print anything
}
func count(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, s)

		time.Sleep(400 * time.Millisecond)
	}
}
