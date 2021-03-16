package main

import "fmt"

func trailingZeroes(n int) int {
	counter := 0
	for i := 5; n/i >= 1; i *= 5 {
		counter += n / i
	}

	return counter
}

func main() {
	fmt.Println(trailingZeroes(15))
}
