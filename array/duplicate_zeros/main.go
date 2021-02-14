package main

import (
	"fmt"
)

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			arr[i] = 0
			i++
		}
	}
}

func duplicateZeros1(arr []int) {
	x := make([]int, len(arr))
	j, len := 0, len(arr)
	for i := 0; i < len; i++ {
		if arr[i] == 0 {
			if j+1 < len {
				x[j+1] = arr[i]
				j += 2
			} else {
				j++
			}
		} else {
			if j < len {
				x[j] = arr[i]
				j++
			}
		}
	}
	for i := 0; i < len; i++ {
		arr[i] = x[i]
	}
}

func main() {
	a1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	a2 := []int{8, 4, 5, 0, 0, 0, 0, 7}
	duplicateZeros(a1)
	duplicateZeros1(a2)

	fmt.Println(a1)
	fmt.Println(a2)
}
