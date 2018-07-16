package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wt sync.WaitGroup
var mutex sync.Mutex
var counter int

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter", counter)
		mutex.Unlock()
	}
	wt.Done()
}

func main() {
	wt.Add(2)
	incrementor("Boo:")
	incrementor("Foo:")
	wt.Wait()
	fmt.Println("Final Counter:", counter)
}
