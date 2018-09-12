package main

import "fmt"

func main() {
	slice := []string{
		"Divya",
		"Eshaan",
		"Aroop",
		"Evaan",
		"Eva",
	}
	slice1 := []string{
		"me",
		"you",
	}
	for i, value := range slice {
		fmt.Println(i, "-", value)
	}
	fmt.Println("-----------------------")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println(slice)
	fmt.Println(append(slice, slice1...))
	fmt.Println(slice[4:])
	fmt.Println(slice[:])
	fmt.Println(slice[2:4])
	//deleting a myString

	slice = append(slice[:3], slice[4:]...)
	fmt.Println(slice)
	t := make([]int, 1, 3)
	t[0] = 1
	t[0]++ //increments 1 to 2
	fmt.Println(t[0])

	//Slice declarations

	a := []string{} // base array is formed but neeed to use append since no length
	var b []string  // no base array , value is "nil" ,need append
	fmt.Println(a, b)

}
