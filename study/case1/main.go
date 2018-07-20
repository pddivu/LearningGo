package main

import "fmt"

type Lamp struct{}

func (l Lamp) on() {
	fmt.Println("On")
}

func (l Lamp) off() {
	fmt.Println("Off")
}

func main() {
	var lamp Lamp

	lamp.on()
	lamp.off()

	Lamp{}.on()
	Lamp{}.off()
}
