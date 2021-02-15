package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	x := make([]int, m+n)
	index1, index2, indexX := 0, 0, 0
	for index1 != m && index2 != n {
		if nums1[index1] < nums2[index2] {
			x[indexX] = nums1[index1]
			index1++

		} else {
			x[indexX] = nums2[index2]
			index2++
		}
		indexX++
	}

	if index1 == m {
		for i := index2; i < n; i++ {
			x[indexX] = nums2[i]
			indexX++
		}
	}

	if index2 == n {
		for i := index1; i < m; i++ {
			x[indexX] = nums1[i]
			indexX++
		}
	}

	copy(nums1, x)
}

func main() {
	a1 := []int{1, 7, 8, 0, 0, 0}
	m1 := 3
	a2 := []int{2, 3, 4}
	n1 := 3
	merge(a1, m1, a2, n1)
	fmt.Println(a1)

	a3 := []int{1, 19, 0, 0, 0, 0}
	m3 := 2
	a4 := []int{2, 3, 4, 20}
	n4 := 4
	merge(a3, m3, a4, n4)
	fmt.Println(a3)
}
