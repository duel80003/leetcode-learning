package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	sum := (1 + len(nums)) * len(nums) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

func main() {
	a1 := []int{8, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(a1))
	a2 := []int{0}
	fmt.Println(missingNumber(a2))
}
