package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}
	num := append(nums1, nums2...)
	fmt.Println(sort.IntsAreSorted(num))
}
