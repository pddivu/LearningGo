package exectest

import "fmt"

var check bool

//Even will check if the entered number is even
//E in the Even has to be capital
//evencheck returns differnet types of arguments
func Even() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)
	//diffrent types of arguments received in order
	d, even := evencheck(n)
	fmt.Println(n, d, even)
}

func evencheck(b int) (int, bool) {
	return b / 2, b%2 == 0
}
