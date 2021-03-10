package main

import "fmt"

func findDuplicate(nums []int) int {
	arr := make([]int, len(nums)-1)
	r := 0
	for i := 0; i < len(nums); i++ {
		if arr[nums[i]-1] != 0 {
			r = arr[nums[i]-1]
			break
		}
		arr[nums[i]-1] = nums[i]
	}
	return r
}

func main() {
	a1 := []int{1, 1, 3, 4, 5, 2}
	fmt.Println(findDuplicate((a1)))

	a2 := []int{1, 1}
	fmt.Println(findDuplicate(a2))
}
