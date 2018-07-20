package exectest

import "fmt"

//Even2 return function
func Even2() {
	var q int
	fmt.Println("Enter a number")
	fmt.Scan(&q)

	evencalc := func(r int) (int, bool) {
		return r / 2, r%2 == 0
	}
	fmt.Println(evencalc(q))

}
