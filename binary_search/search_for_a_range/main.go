package main

import "fmt"

func searchRange(nums []int, target int) []int {
	r := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			if nums[left] == nums[right] {
				return []int{left, right}
			}
			if target != nums[right] {
				right--
			}
			if target != nums[left] {
				left++
			}
		} else if nums[mid] < target {
			if target != nums[left] {
				left = mid + 1
			}
		} else {
			if target != nums[right] {
				right = mid - 1
			}
		}
	}
	return r
}

func main() {
	// [3, 4]
	a1 := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(a1, 8))

	// [1, 4]
	a2 := []int{1, 2, 2, 2, 2, 8}
	fmt.Println(searchRange(a2, 2))

	//[0, 0]
	a3 := []int{1}
	fmt.Println(searchRange(a3, 1))

	// [0, 1]
	a4 := []int{1, 1}
	fmt.Println(searchRange(a4, 1))

	// [-1, -1]
	a5 := []int{}
	fmt.Println(searchRange(a5, 1))
}
