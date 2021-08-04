package main

import "fmt"

var a  = map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	backTrack(digits, 0, "")
	return  res
}

func backTrack(digits string, index int, combination string) {
	if len(digits) == index {
		res = append(res, combination)
	} else {
		temp, _ := a[string(digits[index])]
		for i:= 0; i< len(temp); i++ {
			backTrack(digits, index+1, combination + string(temp[i]))
		}
	}
}

func main() {
	a := "23"
	fmt.Println(letterCombinations(a))
}