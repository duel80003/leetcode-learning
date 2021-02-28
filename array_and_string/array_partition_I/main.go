package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	max := 0
	nums = sortArray(nums)
	for i := 0; i < len(nums); i += 2 {
		t := min(nums[i], nums[i+1])
		max += t
	}
	return max
}

func sortArray(a []int) []int {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	// 4
	s1 := []int{4, 2, 3, 1}
	fmt.Println(arrayPairSum(s1))

	//9
	s2 := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(s2))

}
