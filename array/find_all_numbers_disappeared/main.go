package main

import (
	"fmt"
)

func findDisappearedNumbersExtraSpase(nums []int) []int {
	x := make([]int, len(nums))
	for _, v := range nums {
		x[v-1] = v
	}
	r := make([]int, 0)
	for i, v := range x {
		if v == 0 {
			r = append(r, i+1)
		}
	}
	return r
}

func findDisappearedNumbers(nums []int) []int {
	x := make([]int, len(nums))
	for _, v := range nums {
		x[v-1] = 1
	}
	index := 0
	for i, v := range x {
		if v == 0 {
			x[index] = i + 1
			index++
		}
	}
	return x[:index]
}

func main() {
	a1 := []int{4, 3, 2, 7, 8, 2, 3, 1}

	fmt.Println(findDisappearedNumbers(a1))

	a2 := []int{2, 2, 2}
	fmt.Println(findDisappearedNumbers(a2))
}
