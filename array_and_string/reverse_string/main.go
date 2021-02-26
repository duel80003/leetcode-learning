package main

import (
	"fmt"
)

func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		fmt.Println("aaa", s[start], s[end])
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	fmt.Println(string(s))
}

func main() {
	// olleh
	c1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(c1)

	// ""
	c2 := []byte{}
	reverseString(c2)

	// ""
	c3 := []byte{'h', 'e', 'l', 'l', 'o', '!'}
	reverseString(c3)

}
