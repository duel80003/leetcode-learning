package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	x := make([]int, len(heights))
	r := 0
	copy(x, heights)
	sort.SliceStable(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	for i, v := range x {
		if heights[i] != v {
			r++
		}
	}
	return r
}
func main() {
	a1 := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(a1))

	a2 := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(a2))

	a3 := []int{1, 2, 3, 4, 5}
	fmt.Println(heightChecker(a3))
}
