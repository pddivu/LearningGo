package main

import "fmt"

func main() {
	var test1 = map[int]string{
		1: "Divya",
		2: "Eshaan",
		3: "Aroop",
		4: "Pinku",
	}
	fmt.Println(test1)
	for key, value := range test1 {
		fmt.Println(key, "-", value)
	}
	delete(test1, 4)
	if val, exist := test1[4]; exist {
		fmt.Println(exist, val, "value does exist")
	} else {
		fmt.Println(exist, val, "value does not exist")
	}
}
