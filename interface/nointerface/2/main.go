package main

import "fmt"

type animal struct {
	sound string
}

type cat struct {
	animal
	bite bool
}

type dog struct {
	animal
	bite bool
}

func spec(a interface{}) {
	fmt.Println(a)
}

func main() {
	kittu := dog{animal{"bow"}, true}
	mittu := cat{animal{"meow"}, false}

	spec(mittu)
	spec(kittu)

	pets := []interface{}{kittu, mittu}
	fmt.Println(pets)
}
