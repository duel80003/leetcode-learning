package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	if target < letters[0] || letters[len(letters)-1] <= target {
		return letters[0]
	}
	left, right := 0, len(letters)-1

	for left <= right {
		mid := (right + left) / 2
		if letters[mid] > target {
			right = mid - 1
		} else if letters[mid] <= target {
			left = mid + 1
		}
	}
	return letters[left]
}

// func nextlessLetter(letters []byte, target byte) byte {
// 	if target < letters[0] || letters[len(letters)-1] <= target {
// 		return letters[0]
// 	}
// 	left, right := 0, len(letters)-1

// 	for left <= right {
// 		mid := (right + left) / 2
// 		if letters[mid] >= target {
// 			right = mid - 1
// 		} else if letters[mid] < target {
// 			left = mid + 1
// 		}
// 	}
// 	return letters[left-1]
// }

func main() {

	// e
	b1 := []byte{'b', 'c', 'e', 'f', 'j'}
	fmt.Println(string(nextGreatestLetter(b1, 'd')))

	// c
	b2 := []byte{'c', 'f', 'j'}
	fmt.Println(string(nextGreatestLetter(b2, 'a')))

	// f
	fmt.Println(string(nextGreatestLetter(b2, 'c')))

	// f
	fmt.Println(string(nextGreatestLetter(b2, 'd')))

	// j
	fmt.Println(string(nextGreatestLetter(b2, 'g')))

	// c
	fmt.Println(string(nextGreatestLetter(b2, 'j')))

	// c
	fmt.Println(string(nextGreatestLetter(b2, 'k')))
}
