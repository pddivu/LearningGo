package main

import "fmt"

func main() {
	// There are many ways to represent variadic function & arguments
	// all the methods used here are correct
	data := []float64{34, 45, 65, 55}
	/*avg := average(data...)*/
	avg := average(data)
	fmt.Printf("%T\n", data)

	/*avg := average(34, 45, 65, 55)*/
	fmt.Println(avg)

}

/*func average(num ...float64) float64 {*/
func average(num []float64) float64 {
	fmt.Println(num)
	fmt.Printf("%T\n", num)
	total := 0.0
	for _, i := range num {
		total = total + i
	}
	return total / float64(len(num))
}
