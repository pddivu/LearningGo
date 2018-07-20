package main

import "fmt"

type orange struct {
	quantity int
}

func (o *orange) increment(s int) {
	o.quantity += s
}

func (o *orange) decrement(u int) {
	o.quantity -= u
}

/*func (o *orange) display() string {
	return fmt.Sprintf("%v", o.quantity)
}*/

func main() {
	var fruit orange
	fruit.increment(10)
	fruit.decrement(5)
	fmt.Println(fruit)
}
