package main

import "fmt"

// Even though func world is called first, it will execuete only last
//used for file closing mainly. Files will be closed at the end of execution
func main() {
	defer world()
	hello()

}

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("World")
}
