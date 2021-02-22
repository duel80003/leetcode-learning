package main

import "fmt"

func isValid(s string) bool {
	length := len(s)

	if length == 0 || length%2 != 0 {
		return false
	}

	stack := []byte{}

	for i := range s {
		switch s[i] {
		case '{', '(', '[':
			stack = append(stack, s[i])
		case '}', ')', ']':
			if len(stack) == 0 {
				return false
			}
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] == ']' && x != '[' || s[i] == ')' && x != '(' || s[i] == '}' && x != '{' {
				return false
			}
		default:
			return false
		}

	}

	return len(stack) == 0
}

func main() {
	s1 := "[](){}"
	fmt.Println(isValid(s1))

	s2 := "[](){]"
	fmt.Println(isValid(s2))
}
