package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	table := make(map[string][]string)
	for _, v := range strs {
		b := []byte(v)
		sort.SliceStable(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		table[string(b)] = append(table[string(b)], v)
	}
	// fmt.Println(table)
	r := make([][]string, len(table))
	index := len(table) - 1
	for _, value := range table {
		t := value
		sort.Strings(t)
		r[index] = t
		index--
	}

	return r
}

func main() {

	// [["bat"],["nat","tan"],["ate","eat","tea"]]
	a1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(a1))
}
