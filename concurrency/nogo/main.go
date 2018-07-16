package main

import "fmt"

func main() {
	foo()
	boo()
}

func foo() {
	for i := 1; i < 20; i++ {
		fmt.Println("Foo: ", i)
	}
}
func boo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Boo: ", i)
	}
}
