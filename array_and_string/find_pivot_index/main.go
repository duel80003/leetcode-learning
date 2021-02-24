package main

import "fmt"

func pivotIndex1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left, rigth := 0, 0
		for j := i + 1; j < len(nums); j++ {
			left += nums[j]
		}

		for k := 0; k < i; k++ {
			rigth += nums[k]
		}

		if rigth == left {
			return i
		}
	}
	return -1
}

func pivotIndex(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	left, rigth := 0, 0

	for i := 1; i < len(nums); i++ {
		rigth += nums[i]
	}

	for i := 1; i < len(nums); i++ {
		if left == rigth {
			return i - 1
		}
		left += nums[i-1]
		rigth -= nums[i]
	}
	return -1
}

func main() {
	a1 := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(a1))
	a2 := []int{2, 1, -1}
	fmt.Println(pivotIndex(a2))

	a3 := []int{1, 2, 4, 3}
	fmt.Println(pivotIndex(a3))

}
