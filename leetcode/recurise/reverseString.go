package main

import "fmt"

func reverseStr(str []byte) []byte {
	if str == nil {
		return nil
	}
	if len(str) != 1 {
		return append(reverseStr(str[1:len(str)]), str[0])
	} else {
		return str
	}
}


func reverseStr1(s []byte) {
	if len(s) > 1 {
		s[0] = s[0] + s[len(s)-1]

		s[len(s)-1] = s[0] - s[len(s)-1]
		s[0] = s[0] - s[len(s)-1]
		reverseStr1(s[1:len(s)-2])
	}

}
func main() {
	a:="hellowodezhongguo"
	b:= reverseStr([]byte(a))
    c := 1
    fmt.Println(string(c))
	fmt.Println(string(b))
}