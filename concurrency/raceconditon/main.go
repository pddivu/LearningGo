package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wt sync.WaitGroup
var counter int

func main() {
	wt.Add(2)
	go incrementor("foo:")
	go incrementor("Boo:")
	wt.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		//counter++
		fmt.Println(s, i, "counter-", counter)
	}
	wt.Done()
}
