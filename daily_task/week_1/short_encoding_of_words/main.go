package main

import (
	"fmt"
)

func minimumLengthEncoding(words []string) int {
	table := make(map[string]bool)
	for _, v := range words {
		table[v] = true
	}
	// fmt.Println(table)
	for key := range table {
		for i := 1; i < len(key); i++ {
			s := key[i:]
			delete(table, s)
		}
	}

	sum := 0
	for w := range table {
		sum += len(w) + 1
	}
	return sum
}

func main() {
	// 10
	s1 := []string{"time", "me", "bell"}
	fmt.Println(minimumLengthEncoding(s1))

	// // 9
	// s2 := []string{"time", "me", "me"}
	// fmt.Println(minimumLengthEncoding(s2))

	// // 9
	// s3 := []string{"me", "time"}
	// fmt.Println(minimumLengthEncoding(s3))
	// r := strings.Replace("timez", "me", "", 0)
	// fmt.Println(r)
}
