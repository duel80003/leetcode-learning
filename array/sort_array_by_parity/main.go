package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	s := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[s], nums[i] = nums[i], nums[s]
			s++
		}
	}
	return nums
}
func main() {
	a1 := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(a1))

	a2 := []int{0}
	fmt.Println(sortArrayByParity(a2))

	a3 := []int{2, 3, 1}
	fmt.Println(sortArrayByParity(a3))
}
