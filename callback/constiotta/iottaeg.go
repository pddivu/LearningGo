package main

import "fmt"

const (
	//PI is Pi
	PI = 3.14
	t  = "Learning Go"
)
const (
	A = iota //0
	B = iota //1
	C = iota
)

const (
	//D will be 0
	D = iota
	// E wil be 1
	E
	// F wil be 2
	F
)

const (
	_ = iota
	H = iota * 10
	I = iota * 10
)
const pd = "Divya"

const (
	_ = iota
	//KB is 1024 bites. Shited to the 10th place in the binary chart
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 20)
)

func main() {
	fmt.Println(PI, t, pd)
	fmt.Println(A, B, C)
	fmt.Println(D, E, F)
	fmt.Println(H, I)
	fmt.Printf("%d\t %b\n", KB, KB)
	fmt.Printf("%d\t %b\n", MB, MB)
}
