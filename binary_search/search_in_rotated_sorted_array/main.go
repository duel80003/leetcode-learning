package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}

func main() {
	// 4
	a1 := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(a1, 0))

	// -1
	a2 := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(a2, 3))

	// -1
	a3 := []int{1}
	fmt.Println(search(a3, 0))

	// 0
	a4 := []int{0, 1, 2, 4, 5, 5, 5, 6, 7}
	fmt.Println(search(a4, 6))

	// 1
	a5 := []int{1, 3}
	fmt.Println(search(a5, 3))

	// 0
	a6 := []int{3, 5, 1}
	fmt.Println(search(a6, 3))

	// 1
	a7 := []int{1, 3, 5}
	fmt.Println(search(a7, 3))
}
