package main

import "fmt"

func convertToTitle(n int) string {
	tmp := 0
	result := ""
	for n > 0 {
		tmp = n % 26
		n /= 26
		if tmp == 0 {
			tmp = 26
			n--
		}
		result = string('A'+tmp-1) + result
	}
	return result
}

func main() {
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(28))

	fmt.Println(convertToTitle(2147483647))
	fmt.Println(convertToTitle(701))

}
