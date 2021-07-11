package main

import "fmt"

func fib(n int) int {
	if n==0 || n ==1  {
		return n
	}
	f1 := 0
	f2 := 1
	temp := 0
	for i:=1; i< n; i++ {
		temp = f1 + f2
		f1 = f2
		f2 = temp

	}
	return temp
}

func main() {
	fmt.Println(fib(4))
	str := "烫烫烫烫"
	array := []rune(str)
	n := len(array)
	for i := 0; i < n; i++ {
		ch := array[i]     // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch) //unicode 编码转十进制输出
	}
}