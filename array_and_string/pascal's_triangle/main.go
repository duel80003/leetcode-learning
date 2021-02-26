package main

import "fmt"

func generate(numRows int) [][]int {
	r := make([][]int, numRows)
	max := 2
	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0] = 1
		arr[len(arr)-1] = 1
		r[i] = arr
		if i > 1 {
			length := len(arr) - 1
			mid := length/2 + 1
			arr[mid] = max
			for j := 1; j < mid; j++ {
				c := r[i-1][j-1] + r[i-1][j]
				arr[j] = c
				arr[length-j] = c
				max = arr[mid] + arr[mid-1]
			}
			fmt.Println(max)
		}
	}
	return r
}

func main() {
	fmt.Println(generate(8))
}
