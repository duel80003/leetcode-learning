package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	table := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if index, ok := table[nums[i]]; ok {
			if i-index <= k {
				return true
			}
		}
		table[nums[i]] = i
	}
	return false
}

func main() {
	a1 := []int{1, 2, 3, 1, 3, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(a1, 3))
}
