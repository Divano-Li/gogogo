package main

import (
	"fmt"
	"strconv"
	"strings"
)

var nums []string
var res []int = make([]int, 0)
func printNumbers(n int) []int {
	nums = make([]string, n)
	res = make([]int, 0)

    dfs(0, n)
	return  res
}

func dfs(x int, n int) {
	if x == n {
		 // 可以利用 strconv.Atoi()方法去除前置0
		 /*point := deleteZero(nums)
		 if  point != n {
		 	 temp, _ := strconv.Atoi(strings.Join(nums[point:], ""))
			 res = append(res, temp)
		 }*/
		temp, _ := strconv.Atoi(strings.Join(nums, ""))
		if temp != 0 {
			res = append(res,temp)
		}
	   return
	} else {
		for i:=0; i<10;i++ {
			nums[x] = strconv.Itoa(i)
			dfs(x+1, n)
		}
	}
}

func deleteZero(nums []string) int {
	point := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] == "0" {
			point = point +1
		} else {
			return point
		}
	}
	return point
}


func main() {
	a := printNumbers(2)
	fmt.Println(a)
	fmt.Println(strconv.Atoi("00"))
	b := nil
}