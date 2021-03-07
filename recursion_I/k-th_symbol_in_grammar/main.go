package main

import (
	"fmt"
	"math"
)

func kthGrammar(N int, K int) int {
	fmt.Println(N, K)
	if N == 1 && K == 1 {
		return 0
	}
	mid := int(math.Pow(2, float64(N-2)))
	fmt.Println("mid", mid)
	if K <= mid {
		return kthGrammar(N-1, K)
	}
	if kthGrammar(N-1, K-mid) == 1 {
		fmt.Println("aaa")
		return 0
	}
	return 1
}

func main() {
	// 1
	// fmt.Println(kthGrammar(2, 2))

	//1
	fmt.Println(kthGrammar(4, 6))

	// //0
	// fmt.Println(kthGrammar(1, 1))

	// //0
	// fmt.Println(kthGrammar(30, 434991989))

	// 	// 1
	// 	fmt.Println(kthGrammar(3, 3))
}
