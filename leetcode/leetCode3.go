package main

func lengthOfLongestSubstring(s string) int {
	arr := []byte(s)
	if len(arr) == 0 {
		return 0
	}
	m := map[byte]int{}
	pre := 0
	max := 1
	m[arr[0]] = 0
	for i := 1; i < len(arr); i++ {
		value, ok := m[arr[i]]
		if ok {
			m[arr[i]] = i
			max = maxfun(max, i-pre)
			pre = value + 1
		} else {
			m[arr[i]] = i
			max = maxfun(max, i-pre)
		}
	}
	return max
}

func maxfun(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
