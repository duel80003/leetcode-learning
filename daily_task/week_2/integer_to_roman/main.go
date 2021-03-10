package main

import "fmt"

func intToRoman(num int) string {
	t := 1
	tmp := 0
	p := 5
	r := ""
	pre := ""
	s := "I"
	table := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	for num > 0 {
		n := num % 10 * t
		s = table[t]
		num /= 10
		if v, ok := table[n]; ok {
			r = v + r
		} else {
			if n > p {
				pre = table[p]
				tmp = (n - p) / t
			} else {
				pre = ""
				tmp = n / t
			}
			fmt.Println(tmp)
			for tmp > 0 {
				r = s + r
				tmp--
			}
			r = pre + r
		}
		p *= 10
		t *= 10
	}
	return r
}

func main() {
	fmt.Println(intToRoman(1994))
}
