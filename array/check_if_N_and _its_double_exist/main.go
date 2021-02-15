package main

import "fmt"

func checkIfExist(arr []int) bool {
	hashMap := make(map[int]int)
	for _, v := range arr {
		if _, ok := hashMap[v]; ok {
			return ok
		}
		if _, ok := hashMap[v*4]; ok {
			return ok
		}
		hashMap[v*2] = v
	}
	return false
}

func main() {
	a1 := []int{10, 2, 5, 3}
	fmt.Println(checkIfExist(a1))

	a2 := []int{14, 2, 5, 3, 7}
	fmt.Println(checkIfExist(a2))

	a3 := []int{20, 2, 5, 3}
	fmt.Println(checkIfExist(a3))

	a4 := []int{7, 1, 14, 11}
	fmt.Println(checkIfExist(a4))

	a5 := []int{3, 1, 7, 11}
	fmt.Println(checkIfExist(a5))
}
