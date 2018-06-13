package main

import (
	"fmt"
	"sort"
)

func main() {

	var studygroup []string
	studygroup = []string{"Zeno", "John", "An", "Jenny"}
	sort.Strings(studygroup)
	fmt.Println(studygroup)

	type people []string
	str := people{"Zeno", "John", "An", "Jenny"}
	//stringslice is an interface
	sort.Sort(sort.StringSlice(str))
	//Sort.StringSlice(str).sort()
	fmt.Println(str)

	num := []int{2, 45, 3, 7, 8, 1, 0}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))

	d := sort.IntsAreSorted(num)
	fmt.Println(num, d)

}
