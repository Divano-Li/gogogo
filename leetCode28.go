package main

func strStr(haystack string, needle string) int {
	if needle ==  "" {
		return 0
	}
	a := []byte(haystack)
	b := []byte(needle)

	if len(b) > len(a) {
		return -1
	}
	result := -1
	for i:=0; i<len(a); i++ {
		flag := true
		if len(a[i:len(a)]) < len(b) {
			return result
		} else {
			for j,n := 0,i; j<len(b); j,n =j+1,n+1{
				if a[n] != b[j] {
					flag = false
					break
				}
			}
			if flag {
				result = i
				return result
			}
		}
	}
	return result
}