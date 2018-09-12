package main

import "fmt"

//UTF 8 int 4
func main() {
	for i := 100; i <= 140; i++ {
		fmt.Println(i, "_", string(i), "_", []byte(string(i)))
	}
	// single qoutes 'a' indicates its a rune not a string
	foo := 'a'
	fmt.Println(foo)
	fmt.Println(string(foo))
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", rune(foo))
}
