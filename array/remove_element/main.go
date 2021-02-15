package main

import "fmt"

func removeElement1(nums []int, val int) int {
	count := 0
	lastIndex := len(nums) - 1
	for i := 0; i <= lastIndex; i++ {
		if nums[i] == val {
			count++
			if i < lastIndex {
				tmp := nums[i]
				for nums[lastIndex] == val {
					lastIndex--
					if lastIndex >= 0 {
						count++
					}
					if lastIndex <= i {
						break
					}
				}
				if lastIndex < 0 {
					break
				}
				nums[i] = nums[lastIndex]
				nums[lastIndex] = tmp
				lastIndex--
			} else {
				lastIndex--
			}
		}
	}
	nums = nums[:len(nums)-count]
	return len(nums)
}

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
func main() {
	a1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// fmt.Println(removeElement(a1, 2))
	fmt.Println("5 ==", removeElement(a1, 2))

	a2 := []int{1}
	// fmt.Println(removeElement(a2, 1))
	fmt.Println("0 ==", removeElement(a2, 1))

	a3 := []int{3, 3}
	// fmt.Println(removeElement(a3, 3))
	fmt.Println("0 ==", removeElement(a3, 3))

	a4 := []int{4, 5}
	// fmt.Println(removeElement(a4, 5))
	fmt.Println("1 ==", removeElement(a4, 5))

	a5 := []int{2, 2, 2}
	// fmt.Println(removeElement(a5, 2))
	fmt.Println("0 ==", removeElement(a5, 2))

	a6 := []int{2, 3, 3}
	// fmt.Println(removeElement(a6, 3))
	fmt.Println("1 ==", removeElement(a6, 3))

}
