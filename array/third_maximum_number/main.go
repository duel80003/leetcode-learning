package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	m, s, r := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v != m && v != s {
			if v > m {
				r = s
				s = m
				m = v
			} else if v > s {
				r = s
				s = v
			} else if v > r {
				r = v
			}
		}
	}
	if math.MinInt64 == r {
		return m
	}
	return r
}
func main() {
	a1 := []int{1, 4, 7, 3, 9}
	fmt.Println("4 ==", thirdMax(a1))

	a2 := []int{3, 2, 1}
	fmt.Println("1 ==", thirdMax(a2))

	a3 := []int{1, 1, 1, 1, 2}
	fmt.Println("2 ==", thirdMax(a3))

	a4 := []int{2}
	fmt.Println("2 ==", thirdMax(a4))

	a5 := []int{1, 2, 2, 5, 3, 5}
	fmt.Println("2 ==", thirdMax(a5))

	a6 := []int{2, 2, 2, 1}
	fmt.Println("2 ==", thirdMax(a6))

	a7 := []int{5, 2, 4, 1, 3, 6, 0}
	fmt.Println("4 ==", thirdMax(a7))
}
