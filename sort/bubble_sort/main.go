package main

import "fmt"

func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := bubbleSort([]int{4, 2, 6, 8, 1, 3, 7, 9, 9})
	fmt.Println(arr)
}
