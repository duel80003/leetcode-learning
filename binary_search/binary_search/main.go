package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	a1 := []int{-2, -1, 0, 3, 5, 9, 12, 14, 18}
	fmt.Println(search(a1, 9))
}
