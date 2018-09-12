package main

import (
	"fmt"
	"reflect"
)

func main() {

	var str string
	var length, out int

	fmt.Println(" Enter a string")

	fmt.Scan(&str) /* ----------------> or os.Args */

	/* check if the half of total length substring is same as remaining half.
	   If so for 8 long string, output will be 4, else bring down to 3,2,1 */
	fmt.Println(str)
	if len(str)%2 != 0 {
		length = len(str) - 1
	} else {
		length = len(str)
	}

	half := length / 2

	if reflect.DeepEqual(str[0:half], str[3:6]) == true {
		out = half
	}
	/* compare str [0,3] and str [4,7 ] position ---4*/
	/* compare str [0,2] and str [3,5],[4,6],[5,7] position ---3 */
	/* compare str [0,1] and str [2,3],[3,4],[4,5],[5,6],[6,7] position --2 */
	/*comapre each letter with all the other*/
	fmt.Println(str[0:3], str[3:6], length, half)
	fmt.Println(out)
}
