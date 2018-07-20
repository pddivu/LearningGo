package main

import "fmt"

func main() {

	type sam struct {
		a int
		b string
		c interface{}
	}

	d := sam{a: 1, b: "a", c: "c"}
	e := sam{a: 1, b: "a", c: "c"}

	if d == e {
		fmt.Println("true")
	}

	type slice struct {
		f []int
	}

	g := slice{f: []int{1}}
	fmt.Println(g)
	/* Structure containing slice can not be compared
	h := g

	if h == g {
		fmt.Println("true")
	}*/
}
