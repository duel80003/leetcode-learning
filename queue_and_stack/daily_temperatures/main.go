package main

import "fmt"

func dailyTemperatures(T []int) []int {
	stack := []int{}
	res := make([]int, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			res[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}
	return res
}

func main() {

	// 1 1 0 1 0 0
	a1 := []int{73, 74, 75, 71, 72, 72}
	fmt.Println(dailyTemperatures(a1))

	// 1 1 4 1 2 1 0 0
	a2 := []int{73, 74, 75, 71, 72, 72, 76, 73}
	fmt.Println(dailyTemperatures(a2))

	// 1 1 4 2 1 1 0 0
	a3 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(a3))

}
