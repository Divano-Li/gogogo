package main

import "fmt"
//")()())"
func longestValidParentheses(s string) int {
	left := 0
	right := 0
	max := 0
	for i:=0; i<len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			max = maxInt(max, left*2)
		}
		if right > left {
			left = 0
			right = 0
		}
	}
	left = 0
	right = 0
	for i:=len(s)-1; i>=0; i-- {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			max = maxInt(max, right*2)
		}
		if left > right {
			left = 0
			right = 0
		}
	}
	return max
}

func maxInt(x, y int) int {
	if x>y {
		return x
	} else {
		return y
	}
}

func main() {
	a := "(())"
	fmt.Println(longestValidParentheses(a))
}
