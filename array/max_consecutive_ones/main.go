package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	max, cur := 0, 0
	for _, v := range nums {
		if v == 1 {
			cur++
		} else {
			if max == 0 || cur > max {
				max = cur
			}
			cur = 0
		}
	}
	if cur > max {
		return cur
	}
	return max
}

func main() {
	a1 := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(a1))
	a2 := []int{1, 0, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(a2))
}
