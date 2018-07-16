package main

import (
	"fmt"
	"sync"
	"time"
)

/*func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}*/

var wt sync.WaitGroup

func main() {
	wt.Add(2)
	go foo()
	go boo()
	wt.Wait()
}

func foo() {
	for i := 1; i < 20; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wt.Done()
}
func boo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Boo: ", i)
		time.Sleep(10 * time.Millisecond)
	}
	wt.Done()
}
