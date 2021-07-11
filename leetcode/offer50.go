package main

import "fmt"

func main() {
	s := "你好"
	s1 := []byte(s)
	s2 := []rune(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(byte('a'))
	fmt.Println('A' - 'a')
}

type site struct {
	count int
	index int
}
func firstUniqChar(s string) byte {
	str := []byte(s)
	len := len(str)
	result := " "
	if len == 0 {
		return []byte(result)[0]
	}

	a := make(map[byte]site)
	for i,v := range str {
		value, ok := a[v]
		if ok {
			a[v] = site{value.count + 1, value.index}
		} else {
			a[v] = site{1,i}
		}
	}
	min := site{-1,len}
	for _,v := range a {
		if v.count == 1 && v.index < min.index {
			min = v
		}
	}
	if min.index == len {
		return []byte(result)[0]
	}
    return str[min.index]
}

func firstUniqChar1(s string) byte {
	len := len(s)
	result := " "
	if len == 0 {
		return []byte(result)[0]
	}
    var array  [26]int
	for i:=0 ; i<len; i++ {
		array[s[i] - 'a']++
	}
	for i:=0; i<len;i++ {
		if array[s[i]-'a'] == 1 {
			return s[i]
		}
	}
   return ' '
}