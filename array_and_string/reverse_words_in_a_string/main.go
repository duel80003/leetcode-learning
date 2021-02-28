package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	// fmt.Println(s, len(s))
	r := []byte{}
	// fmt.Println(len(r))
	for i := 0; i < len(s); i++ {
		// fmt.Println(i)
		if s[i] != ' ' {
			r = append(r, s[i])
		}
		if s[i] == ' ' && s[i+1] != ' ' {
			r = append(r, s[i])
		}
	}
	t := strings.Split(string(r), " ")
	last := len(t) - 1
	for i := 0; i < len(t)/2; i++ {
		t[i], t[last] = t[last], t[i]
		last--
	}
	s = strings.Join(t, " ")
	return s
}

func main() {
	fmt.Println(reverseWords("  a good   example  "))
}
