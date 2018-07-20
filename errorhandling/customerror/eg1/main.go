package main

import (
	"errors"
	"fmt"
)

func main() {

	_, err := square(-10)
	if err != nil {
		fmt.Println(err)
	}
}

func square(p float64) (float64, error) {

	if p < 0 {
		return 0, errors.New("The square root of a negative number")
	}

	return 40, nil
}
