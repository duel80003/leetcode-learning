package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) < 0 {
		return
	}
	s, i := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			if nums[i] != nums[s] {
				nums[s] = nums[i]
				nums[i] = 0
			}
			s++
			i++
		} else {
			i++
		}
	}
}

func moveZeroes1(nums []int) {
	x := make([]int, len(nums))
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			x[c] = nums[i]
			c++
		}
	}
	for i, v := range x {
		nums[i] = v
	}
}

func main() {
	a1 := []int{0, 1, 0, 3, 12}

	moveZeroes(a1)
	fmt.Println(a1)

	a2 := []int{1, 0}

	moveZeroes(a2)
	fmt.Println(a2)

	a3 := []int{2, 1}
	moveZeroes(a3)
	fmt.Println(a3)

	a4 := []int{1, 0, 0}
	moveZeroes(a4)
	fmt.Println(a4)

	a5 := []int{1, 0, 1}
	moveZeroes(a5)
	fmt.Println(a5)
}
