package main

import "fmt"

func main() {
	var test1 map[string]string
	fmt.Println(test1 == nil)
	//////////////////////////////
	test2 := make(map[string]string) //or
	//var test2 = make(map[string]string)
	test2["Divya"] = "Divakaran"
	test2["Eshaan"] = "bhagavad"
	fmt.Println(test2)

	///////////////////////////////////
	var test3 = map[string]string{}
	fmt.Println(test3 == nil)
	test3["Divya"] = "Divakaran"
	test3["Eshaan"] = "bhagavad"
	fmt.Println(test3)

	/////////////////////////////////////

	var test4 = map[string]string{
		"Divya":  "Divakaran",
		"Eshaan": "Bhagavad",
	}
	fmt.Println(test4["Eshaan"])
	test4["Aroop"] = "Maliakkal"
	fmt.Println(test4)
	fmt.Println(len(test4))

	delete(test4, "Divya")
	fmt.Println(test4)
}
