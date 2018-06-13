package main

import (
	"fmt"
	"sort"
)

type people struct {
	Name string
	Age  int
}

func (p people) String() string {
	return fmt.Sprintf("Name %s: %d", p.Name, p.Age)
}

type byage []people

func (a byage) Len() int           { return len(a) }
func (a byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byage) Less(i, j int) bool { return a[i].Age < a[j].Age }

//func (a byage) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	list := byage{
		{"Divya", 38},
		{"Eshaan", 7},
		{"Aroop", 38},
	}

	fmt.Println(list[1])
	fmt.Println(list)
	sort.Sort(byage(list))
	fmt.Println(list)
}
