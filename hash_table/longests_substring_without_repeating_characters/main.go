package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	//slicing window left and right pointers
	l := 0
	r := 0
	//define the dict as [128]int much faster than map[byte]int
	dict := [128]int{}
	dict[s[l]]++
	res := 1

	for r < len(s)-1 {
		r++
		dict[s[r]]++
		//if the element appreas more than once, then move left pointer
		for dict[s[r]] > 1 {
			dict[s[l]]--
			l++
		}
		//keep the max length
		res = max(res, r-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 3
	fmt.Println(lengthOfLongestSubstring("abcabcabc"))

	// 3
	fmt.Println(lengthOfLongestSubstring("pwwkew"))

	// 1
	fmt.Println(lengthOfLongestSubstring("bbbbb"))

	//3
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

	//3
	fmt.Println(lengthOfLongestSubstring("dvdf"))

}
