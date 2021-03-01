package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}

func main() {
	a1 := []int{2, 2, 1}
	fmt.Println(singleNumber(a1))
}
