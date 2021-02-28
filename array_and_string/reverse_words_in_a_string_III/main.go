package main

import "fmt"

func reverseWords(s string) string {
	// s, e := 0, 0
	r := []byte{}
	tmp := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			tmp = append(tmp, s[i])
		} else {
			reverseString(tmp)
			tmp = append(tmp, ' ')
			r = append(r, tmp...)
			tmp = tmp[:0]
		}
	}
	reverseString(tmp)
	r = append(r, tmp...)
	return string(r)
}
func reverseString(s []byte) {
	end := len(s) - 1
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[end] = s[end], s[i]
		end--
	}
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	// fmt.Println(reverseWords("Let's take"))

}
