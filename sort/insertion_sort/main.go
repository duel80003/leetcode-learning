package main

import "fmt"

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && key < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
	return nums
}

func main() {
	arr := insertionSort([]int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0})
	fmt.Println(arr)
}
