package main

import "fmt"

func firstUniqChar(s string) int {
	a := make([]int, 26)
	for _, v := range s {
		a[v-'a']++
	}
	for i, v := range s {
		if a[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))

}
