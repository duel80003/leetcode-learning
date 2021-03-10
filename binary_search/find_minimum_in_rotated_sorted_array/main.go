package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	// 0
	a1 := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(a1))

	// 11
	a2 := []int{11, 13, 15, 17}
	fmt.Println(findMin(a2))

	// 1
	a3 := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(a3))
}
