package main

import (
	"errors"
	"fmt"
)

//Errmessage is a error mesage
var Errmessage = errors.New("The square root of a negative number")

func main() {

	_, err := square(-10)
	if err != nil {
		fmt.Println(err)
	}
}

func square(p float64) (float64, error) {

	if p < 0 {
		return 0, Errmessage
	}

	return 40, nil
}
