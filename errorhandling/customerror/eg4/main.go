package main

import (
	"fmt"
	"log"
)

type locationerror struct {
	lat, long string
	err       error
}

func (l locationerror) Error() string {
	return fmt.Sprintf("location errror - %v %v %v ", l.long, l.lat, l.err)
}

func main() {
	_, err := square(-10.11)
	if err != nil {
		log.Println(err)
	}
}

func square(f float64) (float64, error) {
	msg := fmt.Errorf("square root of negative number %v", f)
	return 0, locationerror{"3.44", "5.66", msg}

}
