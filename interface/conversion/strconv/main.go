package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "12"
	y := 6
	z, _ := strconv.Atoi(x)
	fmt.Println(z + y)

	a := 12
	b := "Number of roses "
	c := b + strconv.Itoa(a)
	fmt.Println(c)
	fmt.Println(b + strconv.Itoa(a))

	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	g, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(g, f, i, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	s := strconv.FormatBool(true)
	r := strconv.FormatFloat(3.1415, 'E', -1, 64)
	q := strconv.FormatInt(-42, 16)
	p := strconv.FormatUint(42, 16)

	fmt.Println(s, r, q, p)
}
