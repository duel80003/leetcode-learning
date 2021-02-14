package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}

func sortedSquares1(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	l, r := 0, length-1
	index := r
	for l <= r {
		fmt.Printf("l: %v, r: %v \n", l, r)
		if nums[l]+nums[r] < 0 {
			ans[index] = nums[l] * nums[l]
			l++
		} else {
			ans[index] = nums[r] * nums[r]
			r--
		}
		index--
		fmt.Println(ans)
	}
	return ans
}

func main() {
	a1 := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares1(a1))

	// a2 := []int{-7, -3, 2, 3, 11}
	// fmt.Println(sortedSquares1(a2))

}
