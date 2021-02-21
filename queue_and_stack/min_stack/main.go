package main

import (
	"fmt"
	"math"
)

// MinStack stack implemention
type MinStack struct {
	Nums    []int
	MinNums []int
}

// Constructor init MinStack
func Constructor() MinStack {
	nums, minNums := make([]int, 0, 5), make([]int, 0, 5)
	minNums = append(minNums, math.MaxInt32)
	m := MinStack{nums, minNums}
	return m
}

// Push new value to top
func (s *MinStack) Push(x int) {
	s.Nums = append(s.Nums, x)
	if x <= s.MinNums[len(s.MinNums)-1] {
		s.MinNums = append(s.MinNums, x)
	}
}

// Pop remove value from top
func (s *MinStack) Pop() {
	if s.Nums[len(s.Nums)-1] == s.MinNums[len(s.MinNums)-1] {
		s.MinNums = s.MinNums[:len(s.MinNums)-1]
	}
	s.Nums = s.Nums[:len(s.Nums)-1]

}

// Top return top value
func (s *MinStack) Top() int {
	return s.Nums[len(s.Nums)-1]
}

// GetMin return the minimum value in stack
func (s *MinStack) GetMin() int {
	return s.MinNums[len(s.MinNums)-1]
}

func main() {
	// s1 := Constructor()
	// s1.Push(-2)
	// s1.Push(0)
	// s1.Push(-3)
	// // fmt.Println(s1.Nums)
	// fmt.Println(s1.GetMin())
	// s1.Pop()
	// fmt.Println(s1.Top())
	// fmt.Println(s1.GetMin())

	// s2 := Constructor()
	// s2.Push(-2)
	// fmt.Println(s2.GetMin())

	// s3 := Constructor()
	// s3.Push(-2)
	// s3.Push(0)
	// s3.Push(-3)
	// s3.Push(6)
	// s3.Pop()
	// s3.Pop()
	// s3.Pop()
	// // s3.Pop()
	// fmt.Println(s3.Top())
	// fmt.Println(s3.GetMin())

	// s4 := Constructor()
	// s4.Push(2)
	// s4.Push(0)
	// s4.Push(3)
	// s4.Push(0)
	// fmt.Println(s4.GetMin())
	// s4.Pop()
	// fmt.Println(s4.GetMin())

	// 2147483647, 2147483646, 2147483646, 2147483647, 2147483647
	s5 := Constructor()
	s5.Push(math.MaxInt32 - 1)
	s5.Push(math.MaxInt32 - 1)
	s5.Push(math.MaxInt32)
	fmt.Println(s5.Top())
	s5.Pop()
	fmt.Println(s5.GetMin())
	s5.Pop()
	fmt.Println(s5.GetMin())
	s5.Pop()
	s5.Push(math.MaxInt32)
	// fmt.Println(s5)
	fmt.Println(s5.Top()) //
	fmt.Println(s5.GetMin())
	s5.Push(math.MinInt32)
	fmt.Println(s5.Top())
	fmt.Println(s5.GetMin())
	s5.Pop()
	fmt.Println(s5.GetMin()) //
}
