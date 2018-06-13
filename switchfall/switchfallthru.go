package main

import "fmt"

func main() {
	var mark int
	fmt.Println("Enter your score")
	fmt.Scan(&mark)
	fmt.Printf("%d\n", mark)
	fmt.Println("", mark)

	switch {
	case mark > 75:
		fmt.Println("You have score A grade")
	case mark > 50:
		fmt.Println("You have score B grade")
	default:
		fmt.Println("You have score C grade")
	}
}
