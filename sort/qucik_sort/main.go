package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	for i, v := range a {
		if v < a[right] {
			a[left], a[i] = v, a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	fmt.Println(a, left)

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {

	slice := generateSlice(10)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}
