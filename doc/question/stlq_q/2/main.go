package main

import "fmt"

func main() {
	var nums1 []interface{}
	nums2 := []int{1, 2, 3}
	num3 := append(nums1, nums2)
	fmt.Println(len(num3))
	// What is output here ?
	// A. 1
	// B. 3
	// C. 4
	// D. panic
}
