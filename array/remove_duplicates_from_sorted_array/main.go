package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count, pointer := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != pointer {
			nums[count], pointer = nums[i], nums[i]
			count++
		}

	}
	return count
}

func main() {

	a1 := []int{0, 1, 2, 2, 2, 3, 4, 5, 5}
	fmt.Println("6 ==", removeDuplicates(a1))
	fmt.Println(a1)

	a2 := []int{1}
	fmt.Println("1 ==", removeDuplicates(a2))
	fmt.Println(a2)

	a3 := []int{3, 3}
	fmt.Println("1 ==", removeDuplicates(a3))
	fmt.Println(a3)

	a4 := []int{4, 5}
	fmt.Println("2 ==", removeDuplicates(a4))
	fmt.Println(a4)

	a5 := []int{2, 2, 2}
	fmt.Println("1 ==", removeDuplicates(a5))
	fmt.Println(a5)

	a6 := []int{2, 3, 3}
	fmt.Println("2 ==", removeDuplicates(a6))
	fmt.Println(a6)

	a7 := []int{1, 1, 2}
	fmt.Println("2 ==", removeDuplicates(a7))
	fmt.Println(a7)

	a8 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("5 ==", removeDuplicates(a8))
	fmt.Println(a8)
}
