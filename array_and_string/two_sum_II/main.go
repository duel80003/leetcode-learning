package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		v, ok := table[numbers[i]]
		table[target-numbers[i]] = i + 1
		if ok {
			return []int{v, i + 1}
		}
	}
	return []int{}
}

func main() {
	a1 := []int{2, 7, 11, 15}
	fmt.Println(twoSum(a1, 9))

	a2 := []int{2, 3, 4}
	fmt.Println(twoSum(a2, 6))
}
