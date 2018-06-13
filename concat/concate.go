package main

import "fmt"

var f, m, l string

func main() {
	fmt.Println("Enter your first name")
	fmt.Scan(&f)
	fmt.Println("Enter your middle name")
	fmt.Scan(&m)
	fmt.Println("Enter your Last name")
	fmt.Scan(&l)
	fmt.Println(fullname(f, m, l))
	fmt.Println(l + m + f)
}

//Sprint will concatenate strings with out space
func fullname(f, m, l string) (string, string) {
	return fmt.Sprint(f, " ", m, " ", l), fmt.Sprintf(l, " ", m, " ", f)
}
