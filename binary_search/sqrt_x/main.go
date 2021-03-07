package main

import "fmt"

func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	left, right, mid, result := 0, x, 0, 0
	for left <= right {
		mid = (left + right) / 2
		if mid*mid <= x {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}

	}
	return result
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(7))
	fmt.Println(mySqrt(25))
}
