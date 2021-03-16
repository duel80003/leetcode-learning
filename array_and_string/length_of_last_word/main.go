package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	arr := strings.Fields(strings.Trim(s, " "))
	// fmt.Println(len(arr))
	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
}
