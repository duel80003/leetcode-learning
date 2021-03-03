package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	table := make(map[int]int)
	for _, v := range nums {
		table[v]++
	}
	index := 0
	missing := 0
	duplication := 0
	for {
		if v, ok := table[index]; ok {
			if v == 2 {
				duplication = index
			}
		} else {
			missing = index
		}
		index++
		if missing != 0 && duplication != 0 {
			return []int{duplication, missing}
		}
	}
}

func main() {
	// [2, 1]
	a1 := []int{2, 2}
	fmt.Println(findErrorNums(a1))

	// [2, 3]
	a2 := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(a2))

	// [1, 2]
	a3 := []int{1, 1}
	fmt.Println(findErrorNums(a3))

	// [2, 1]
	a4 := []int{3, 2, 2}
	fmt.Println(findErrorNums(a4))
}
