package main

import "fmt"

func replaceElements(arr []int) []int {
	m := -1
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = m
		if tmp > m {
			m = tmp
		}
	}
	return arr
}

func main() {
	a1 := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(a1))

	a2 := []int{0, 1, 2, 2, 5, 4, 2}
	fmt.Println(replaceElements(a2))

	a3 := []int{400}
	fmt.Println(replaceElements(a3))

}
