package main

import "fmt"

func dominantIndex(nums []int) int {
	maxNumIndex := 0
	for i, v := range nums {
		if v > nums[maxNumIndex] {
			maxNumIndex = i
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[maxNumIndex] && nums[maxNumIndex] < nums[i]*2 {
			return -1
		}
	}
	return maxNumIndex
}

func main() {
	a1 := []int{1, 2, 3, 4}
	fmt.Println(dominantIndex(a1))

	a2 := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(a2))
}
