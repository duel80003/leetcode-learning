package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	table := make(map[int]int)
	r := []int{}
	for _, v := range nums1 {
		table[v]++
	}
	for _, v := range nums2 {
		if val, ok := table[v]; ok && val > 0 {
			r = append(r, v)
			table[v]--

		}
	}
	return r
}

func main() {
	//[2]
	a1 := []int{1, 2, 2, 1}
	a2 := []int{2, 2}
	fmt.Println(intersect(a1, a2))

	//[9, 4]
	a3 := []int{4, 9, 5}
	a4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(a3, a4))
	// [1]
	a5 := []int{3, 1, 2}
	a6 := []int{1, 1}
	fmt.Println(intersect(a5, a6))
}
