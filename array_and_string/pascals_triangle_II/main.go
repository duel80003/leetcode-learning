package main

import "fmt"

func getRow(rowIndex int) []int {
	r := make([][]int, rowIndex+1)
	max := 2
	for i := 0; i < rowIndex+1; i++ {
		arr := make([]int, i+1)
		arr[0] = 1
		arr[len(arr)-1] = 1
		r[i] = arr
		if i > 1 {
			length := len(arr) - 1
			mid := (length + 1) / 2
			arr[mid] = max
			for j := 1; j < mid; j++ {
				c := r[i-1][j-1] + r[i-1][j]
				arr[j] = c
				arr[length-j] = c
				max = arr[mid] + arr[mid-1]
			}
		}
	}
	return r[len(r)-1]
}

func main() {
	fmt.Println(getRow(3))
}
