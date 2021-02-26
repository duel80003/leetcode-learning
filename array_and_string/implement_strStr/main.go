package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	// needleLen := len(needle)
	var buffer strings.Builder
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			break
		}
		if haystack[i] == needle[0] {
			tmp := haystack[i : i+len(needle)]
			buffer.WriteString(tmp)
			if needle == buffer.String() {
				return i
			}
			buffer.Reset()
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))           //2
	fmt.Println(strStr("aaaaa", "bba"))          // -1
	fmt.Println(strStr("", ""))                  //0
	fmt.Println(strStr("aaa", "aaaa"))           // -1
	fmt.Println(strStr("mississippi", "issip"))  //4
	fmt.Println(strStr("mississippi", "issipi")) //4
}
