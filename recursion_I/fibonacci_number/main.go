package main

import "fmt"

var cache map[int]int = make(map[int]int)

func fib(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	r := 0
	if n < 2 {
		r = n
	} else {
		r = fib(n-1) + fib(n-2)
	}
	cache[n] = r
	return r
}

func main() {
	fmt.Println(fib(8))
}
