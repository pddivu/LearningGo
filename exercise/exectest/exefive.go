package exectest

import "fmt"

// Foo can have differnt count of arguments
func Foo() {
	foo()
	foo(1, 2, 3)
	foo(1, 2)
	slice := []int{1, 2, 3, 4}
	foo(slice...)
}
func foo(data ...int) {
	fmt.Println(data)
}
