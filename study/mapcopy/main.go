package main

import "fmt"

func main() {

	a := map[string]bool{"A": true, "B": false}
	b := make(map[string]bool)

	//can copy the content by this step too
	//b = a
	//fmt.Println(a, b)
	for key, value := range a {
		b[key] = value
	}

	fmt.Println(a, b)
}
