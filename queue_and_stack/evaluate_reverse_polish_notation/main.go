package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	if len(tokens) == 1 {
		v, _ := strconv.Atoi(tokens[0])
		return v
	}
	r := 0
	for i := 0; i < len(tokens); i++ {
		v, e := strconv.Atoi(tokens[i])
		if e == nil {
			stack = append(stack, v)
		} else {
			v1 := stack[len(stack)-1]
			v2 := stack[len(stack)-2]
			switch tokens[i] {
			case "+":
				r = v2 + v1
			case "-":
				r = v2 - v1
			case "*":
				r = v2 * v1
			case "/":
				r = v2 / v1
			}
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = r
		}
	}
	return r
}
func main() {
	a1 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(a1))

	a2 := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(a2))

	a3 := []string{"18"}
	fmt.Println(evalRPN(a3))

}
