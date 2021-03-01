package main

import "fmt"

func containsDuplicate(nums []int) bool {
	table := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if v, ok := table[nums[i]]; ok {
			return v
		}
		table[nums[i]] = true
	}
	return false
}

func main() {
	a1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(a1))
}
