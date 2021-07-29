package main

import (
	"fmt"
)

func maxLenHuiwenStr1(s string) string {
	if len(s) < 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}

	max := 1
	begin := 0
	for i := 0; i < len(s); i++ {
		left, right := expend(i, i, s)
		if max < right-left+1 {
			max = right - left + 1
			begin = left
		}
		if i+1 < len(s) && s[i] == s[i+1] {
			left1, right1 := expend(i, i+1, s)
			if max < right1-left1+1 {
				max = right1 - left1 + 1
				begin = left1
			}
		}
	}
	return s[begin : begin+max]
}

func expend(i, j int, s string) (int, int) {
	if i < 0 || j >= len(s) || s[i] != s[j] {
		return i + 1, j - 1
	} else {
		return expend(i-1, j+1, s)
	}
}

func maxLenHuiwenStr(s string) string {
	if len(s) < 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}
	//temp := map[string]bool{}
	l := len(s)
	temp1 := [][]bool{}
	for i := 0; i < len(s); i++ {
		temp1 = append(temp1, make([]bool, l))
	}
	for i := 0; i < len(s); i++ {
		temp1[i][i] = true
	}
	max := 1
	begin := 0
	for L := 2; L <= len(s); L++ {
		for i := 0; i < len(s)+1-L; i++ {
			j := L + i - 1
			if s[i] == s[j] {
				if j-i < 3 {
					temp1[i][j] = true
				} else {
					temp1[i][j] = temp1[i+1][j-1]
				}
			} else {
				temp1[i][j] = false
			}
			if temp1[i][j] && L > max {
				max = L
				begin = i
			}
			/*if j-i == 1 {
				if s[j] == s[i] {
					temp[strconv.Itoa(i) + strconv.Itoa(j)] = true
					if L > max {
						max = L
						begin = i
					}
				} else {
					temp[strconv.Itoa(i) + strconv.Itoa(j)] = false
				}
			} else {
				index :=  strconv.Itoa(i+1) + strconv.Itoa(j-1)
				if temp[index] {
					if s[j] == s[i] {
						temp[strconv.Itoa(i) + strconv.Itoa(j)] = true
						if L > max {
							max = L
							begin = i
						}
					} else {
						temp[strconv.Itoa(i) + strconv.Itoa(j)] = false
					}
				} else {
					temp[strconv.Itoa(i) + strconv.Itoa(j)] = false
				}
			}*/
		}
	}
	return s[begin : begin+max]
}

func main() {
	a := "abba"
	fmt.Println(a[0:1])
	fmt.Println(maxLenHuiwenStr(a))
}
