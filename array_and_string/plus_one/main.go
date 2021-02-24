package main

import "fmt"

func plusOne(digits []int) []int {
	add := 0
	index := len(digits) - 1
	digits[index]++
	if digits[index] == 10 {
		for index+1 > 0 {
			digits[index] = digits[index] + add
			if digits[index] >= 10 {
				add = 1
				digits[index] = 0
			} else {
				add = 0
				break
			}
			index--
		}
	}
	if add != 0 {
		return append([]int{add}, digits...)
	}
	return digits
}

func main() {
	a1 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(a1))
	a2 := []int{9, 9}
	fmt.Println(plusOne(a2))
	a3 := []int{0}
	fmt.Println(plusOne(a3))
	a4 := []int{8, 9, 9, 9}
	fmt.Println(plusOne(a4))
}
