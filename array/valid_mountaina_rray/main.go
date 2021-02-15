package main

import "fmt"

func validMountainArray1(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	m, c, goTopStep := -1, -1, 0
	hasTop := false
	for _, v := range arr {
		// fmt.Printf("m: %v, c: %v, v: %v, hasTop: %v \n", m, c, v, hasTop)
		if v == c {
			return false
		}
		if (v > c || goTopStep <= 1) && hasTop {
			return false
		}
		c = v
		if v > m {
			m = v
			goTopStep++
		} else {
			hasTop = true
		}
	}
	return hasTop
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	goDown := false
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if goDown {
			if arr[i] > arr[i-1] {
				return false
			}
		} else {
			if arr[i] < arr[i-1] {
				if i == 1 {
					return false
				}
				goDown = true
			}
		}
	}

	return goDown
}

func main() {
	a1 := []int{0, 1, 2, 3, 5, 4, 2}
	fmt.Println("true ==", validMountainArray(a1))

	a2 := []int{0, 1, 2, 2, 5, 4, 2}
	fmt.Println("false ==", validMountainArray(a2))

	a3 := []int{2, 1}
	fmt.Println("false ==", validMountainArray(a3))

	a4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("false ==", validMountainArray(a4))

	a5 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("false ==", validMountainArray(a5))
}
