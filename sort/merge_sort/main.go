package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice))
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// func mergeSort(items []int) []int {
// 	if len(items) == 1 {
// 		return items
// 	}
// 	middle := len(items) / 2
// 	return merge(mergeSort(items[:middle]), mergeSort(items[middle:]))
// }

// func merge(left, right []int) (result []int) {
// 	size, lIndex, rIndex := len(left)+len(right), 0, 0
// 	result = make([]int, size)
// 	for i := 0; i < size; i++ {
// 		if lIndex >= len(left) && rIndex < len(right) {
// 			result[i] = right[rIndex]
// 			rIndex++
// 		} else if rIndex >= len(right) && lIndex < len(left) {
// 			result[i] = left[lIndex]
// 			lIndex++
// 		} else if left[lIndex] < right[rIndex] {
// 			result[i] = left[lIndex]
// 			lIndex++
// 		} else {
// 			result[i] = right[rIndex]
// 			rIndex++
// 		}
// 	}
// 	return result
// }

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(left, right []int) []int {
	size, lIndex, rIndex := len(left)+len(right), 0, 0
	result := make([]int, size)

	for i := 0; i < size; i++ {
		if lIndex >= len(left) && rIndex < len(right) {
			result[i] = right[rIndex]
			rIndex++
		} else if rIndex >= len(right) && lIndex < len(left) {
			result[i] = left[lIndex]
			lIndex++
		} else if left[lIndex] < right[rIndex] {
			result[i] = left[lIndex]
			lIndex++
		} else {
			result[i] = right[rIndex]
			rIndex++
		}
	}
	return result
}
