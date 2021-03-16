package main

import "fmt"

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}

	sieve := []bool{}
	sieve = append(sieve, []bool{false, false}...)
	for i := 2; i < n; i++ {
		sieve = append(sieve, true)
	}

	count := 0

	// create Sieve of Eratosthenes
	for i := 2; i*i < n; i++ {
		if !sieve[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			sieve[j] = false
		}
	}
	// calculate count of primes number
	for i := 2; i < n; i++ {
		if sieve[i] {
			count++
		}
	}
	return count
}

func main() {
	// fmt.Println(countPrimes(10))
	n := 4
	if n > 0 {
		fmt.Println(n&(n-1) == 0)
	}
}
