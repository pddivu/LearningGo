package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	var str2 string
	fmt.Println("Enter a string")
	fmt.Scan(&str2)

	var str string
	str = "abbc"

	check := anagram(str, str2)
	if check {
		fmt.Println("The words are anagram")
	} else {
		fmt.Println("The words are not anagram")
	}

}

func anagram(st1, st2 string) bool {
	var chk bool
	var st3, st4 []string
	if len(st1) == len(st2) {
		for i := 0; i < len(st1); i++ {
			st3 = append(st3, string(st1[i]))
			st4 = append(st4, string(st2[i]))
		}

		sort.Strings(st3)
		sort.Strings(st4)
		fmt.Println(st3, st4)
		chk = reflect.DeepEqual(st3, st4)
	} else {
		chk = false
	}

	return chk

}
