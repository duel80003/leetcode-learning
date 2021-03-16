package main

func majorityElement(nums []int) int {
	majority, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if majority == nums[i] {
			count++
		} else {
			if count <= 0 {
				majority = nums[i]
				count++
			} else {
				count--
			}
		}
	}
	return majority
}

func main() {

}
