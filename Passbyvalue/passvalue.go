package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m)
	changeme(m)
	fmt.Println(m)

	n := make(map[string]int)
	change(n)
	fmt.Println(n["Todd"])
}

func changeme(a []string) {
	a[0] = "Divya"
	fmt.Println(a)
}

func change(b map[string]int) {
	b["Todd"] = 44
}
