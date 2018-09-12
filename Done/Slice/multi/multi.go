package main

import "fmt"

func main() {
	report := make([][]string, 0)
	student := make([]string, 3)

	student[0] = "Divya"
	student[1] = "Divakaran"
	student[2] = "78%"
	report = append(report, student)

	student1 := make([]string, 3)
	student1[0] = "Eshaan"
	student1[1] = "Maliakkal"
	student1[2] = "88%"

	report = append(report, student1)
	fmt.Println(report)

	lists := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {

		list := make([]int, 0)
		for j := 0; j < 4; j++ {
			list = append(list, j)

		}
		lists = append(lists, list)
	}
	fmt.Println(lists)
}
