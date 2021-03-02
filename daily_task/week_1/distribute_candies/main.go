package main

import "fmt"

func distributeCandies(candyType []int) int {
	t := len(candyType) / 2
	table := make(map[int]bool)
	for _, v := range candyType {
		table[v] = true
		if len(table) > t {
			return t
		}
	}
	return len(table)
}

func main() {
	// 3
	a1 := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(a1))
	// 2
	a2 := []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(a2))
}
