package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	result, level, t := 0, 0, 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += t * int(columnTitle[i]-64)
		level++
		t *= 26
	}
	return result
}

func main() {
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
	fmt.Println(titleToNumber("RAA"))
	fmt.Println(titleToNumber("FXSHRXW"))
	// fmt.Println('R') // 936
	// fmt.Println('X') // 624
	// fmt.Println('W') // 23  //1583
	t := "0011"
	fmt.Println(t[2:4])
}
