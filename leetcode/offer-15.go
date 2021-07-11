package main

import "fmt"

func main() {

}
// 10 进制 换 2 进制的做法
func hammingWeight(num uint32) int {
	count := 0
	var a uint32 = 2
	for  num!=0  {
		a = num % 2
		num = num / 2
		if a != 0 {
			count++
		}
	}
	return count
}

// 位运算做法

func hammingWeight1(num uint32) int {
	var res uint32 =  0
	count := 0
	for num != 0 {
		res = num &1
		num = num >> 1
		if res !=0 {
			count++
		}
	}
	return count
}