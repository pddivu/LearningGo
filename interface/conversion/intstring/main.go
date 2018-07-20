package main

import "fmt"

func main() {
	x := 2
	y := 2.34
	z := x + int(y)
	fmt.Println(z)

	a := float64(x) + y
	fmt.Println(a)

	var b rune
	b = 'a'
	var c int32
	c = 'c'
	fmt.Println(b, c)
	fmt.Println(string(b), string(c))

	fmt.Println(string([]byte{'a', 'b', 'c', 'd'}))
	fmt.Println([]byte("hello"))
}
