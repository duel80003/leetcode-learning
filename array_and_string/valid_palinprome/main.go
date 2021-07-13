package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	s = strings.ToLower(s)
	for left <= right {
		iaValidLeft := isValidChar(s[left])
		isValidRight := isValidChar(s[right])
		if iaValidLeft && isValidRight {
			if s[left] != s[right] {
				return false
			}
		}
		if iaValidLeft && !isValidRight {
			right--
		} else if !iaValidLeft && isValidRight {
			left++
		} else {
			right--
			left++
		}
	}
	return true
}

func isValidChar(b byte) bool {
	return b >= 'a' && b <= 'z'
}
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println(isValidChar(' '))
}
