package main

import (
	"fmt"
	"reflect"
)

func main() {

	fruit := map[string]int{
		"Orange": 200,
		"Apple":  150,
		"Grape":  70,
	}

	//var keys []string
	//var keys []int
	/*
		for key := range fruit {
			keys = append(keys, key)
		}*/
	keys := reflect.ValueOf(fruit).MapKeys()
	fmt.Println(keys)
	/*
		for _, v := range keys {
			fmt.Printf("%s : %v\n", v, fruit[v])
		}*/
}
