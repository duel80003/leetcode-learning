package main

import (
	"fmt"
)

func findNumbers(nums []int) int {
	result := 0
	for _, v := range nums {
		count := 0
		for v != 0 {
			v /= 10
			count++
		}
		if count%2 == 0 {
			result++
		}
	}
	return result
}

func main() {
	a1 := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(a1))

	a2 := []int{555, 901, 482, 1771}
	fmt.Println(findNumbers(a2))

}
