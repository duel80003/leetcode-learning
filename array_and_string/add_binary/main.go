package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	len1 := len(a) - 1
	len2 := len(b) - 1
	add := 0
	t := 0
	r := []byte{}
	for len1 >= 0 || len2 >= 0 {
		t += add
		if len1 >= 0 {
			t += coverByte(a[len1])
		}
		if len2 >= 0 {
			t += coverByte(b[len2])
		}
		// fmt.Println("t", t)
		if t == 3 {
			add = 1
			r = append(r, 49)
		} else if t == 2 {
			add = 1
			r = append(r, 48)
		} else if t == 1 {
			add = 0
			r = append(r, 49)
		} else {
			add = 0
			r = append(r, 48)
		}
		// fmt.Println("add", add)
		t = 0
		len1--
		len2--
	}
	// fmt.Println(r, add)

	if add == 1 {
		r = append(r, '1')
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	// fmt.Println(r)
	return string(r)
}

func coverByte(b byte) int {
	if b == '1' {
		return 1
	}
	return 0
}

func main() {
	// 1100
	fmt.Println(addBinary("1001", "11"))
	// 11000
	fmt.Println(addBinary("1101", "11"))

	// 10101
	fmt.Println(addBinary("1010", "1011"))
}
