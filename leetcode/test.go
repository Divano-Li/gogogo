package main

import "fmt"

func maxSubStrLen(s string) int {
	a  := []byte(s)
	slow:=0
	b := map[byte]int{a[0]:0}
	max := 1
	i := 1
	for ;i<len(a);i++ {
		v, ok := b[a[i]]
		if ok && v >= slow {
			max = maxNum(max, i-slow)
			slow = b[a[i]]+1
			b[a[i]] = i
		} else {
			b[a[i]] = i
		}
	}
	return maxNum(max, i-slow)
}

func maxNum(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main(){
	s := "aabbvv"
	fmt.Println(maxSubStrLen(s))
}

