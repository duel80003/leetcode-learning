package main

import "fmt"

const badVersion = 4

func isBadVersion(n int) bool {
	return n == badVersion
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		fmt.Println(mid)
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(firstBadVersion(5))
}
