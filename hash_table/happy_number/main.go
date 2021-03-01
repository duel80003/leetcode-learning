package main

import "fmt"

func isHappy(n int) bool {
	table := make(map[int]bool)
	for {
		sum, x := 0, n
		for x > 0 {
			t := x % 10
			sum += t * t
			x /= 10
		}
		if _, ok := table[sum]; ok {
			break
		}
		n = sum
		table[n] = true
	}
	return n == 1
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(1111111))
}
