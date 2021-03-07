package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	halfN := n / 2
	product := pow(x, halfN)
	if n%2 == 0 {
		return product * product
	}
	return product * product * x
}

func main() {
	fmt.Println(myPow(2.0000, 3))
}
