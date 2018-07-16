package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wt sync.WaitGroup
var counter int64

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		//fmt.Println(s, i, "Counter", atomic.LoadInt64(&counter))
		fmt.Println(s, i, "Counter", counter)
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
