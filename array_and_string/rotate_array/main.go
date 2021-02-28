package main

import "fmt"

func rotate(nums []int, k int) {
	last := len(nums)
	for i := 0; i < len(nums)/2; i++ {
		last--
		nums[i], nums[last] = nums[last], nums[i]
	}
	last = k % len(nums)
	for i := 0; i < (k%len(nums))/2; i++ {
		last--
		nums[i], nums[last] = nums[last], nums[i]
	}
	last = len(nums)
	for i, j := k%len(nums), 0; j < (len(nums)-(k%len(nums)))/2; i, j = i+1, j+1 {
		last--
		nums[i], nums[last] = nums[last], nums[i]
	}
	fmt.Println(nums)
}

func main() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a1, 3)
	a2 := []int{1, 2}
	rotate(a2, 3)
	a3 := []int{1, 2, 3}
	rotate(a3, 4)
	a4 := []int{1, 2, 3, 4, 5, 6}
	rotate(a4, 11)
}
