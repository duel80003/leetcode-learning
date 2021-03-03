package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	table := make(map[byte]bool)

	for i := 0; i < len(jewels); i++ {
		table[jewels[i]] = true
	}

	r := 0

	for i := 0; i < len(stones); i++ {
		if _, ok := table[stones[i]]; ok {
			r++
		}
	}
	return r
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))

}
