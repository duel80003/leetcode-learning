package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	s1, s2 := 1, 2
	for i := 0; i < n-2; i++ {
		s1, s2 = s2, s1+s2
	}
	return s2
}

func main() {
	fmt.Println(climbStairs(3))
}
