package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	table := make(map[int]int)
	r := []int{}

	for _, v := range nums1 {
		table[v] = 1
	}

	for _, v := range nums2 {
		if _, ok := table[v]; ok {
			r = append(r, v)
			delete(table, v)
		}
	}

	return r
}

func main() {
	//[2]
	a1 := []int{1, 2, 2, 1}
	a2 := []int{2, 2}
	fmt.Println(intersection(a1, a2))

	//[9, 4]
	a3 := []int{4, 9, 5}
	a4 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(a3, a4))
}
