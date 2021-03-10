package main

import "fmt"

func reverse(s string) string {
	left, right := 0, len(s)-1
	r := []byte(s)
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}
func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	} else if reverse(s) == s {
		return 1
	} else {
		return 2
	}
}

func main() {
	fmt.Println(removePalindromeSub("ababa"))
	fmt.Println(removePalindromeSub("abb"))
}
