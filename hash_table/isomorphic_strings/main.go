package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	table1 := make(map[byte]byte)
	table2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		v1, v1ok := table1[s[i]]
		_, v2ok := table2[t[i]]
		if !(v1ok || v2ok) {
			table1[s[i]] = t[i]
			table2[t[i]] = s[i]
		} else if v1 != t[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
}
