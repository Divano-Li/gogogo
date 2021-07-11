package main

import "strings"

func replaceSpace(s string) string {
	count := 0
	temp := []rune(s)

	for i:=0 ;i< len(temp); i++ {
		if temp[i] == ' ' {
			count++
		}
	}
	strings.ReplaceAll()
	temp1 := make([]rune, count*2 + len(temp))
}