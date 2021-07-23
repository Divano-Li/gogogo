package main

import (
	"fmt"
	"sort"
)

//暴力法
func isCovered(ranges [][]int, left int, right int) bool {
	a := map[int]int{}
	for i := 0; i < len(ranges); i++ {
		for j := ranges[i][0]; j <= ranges[i][1]; j++ {
			a[j] = 1
		}
	}
	for j := left; j <= right; j++ {
		_, ok := a[j]
		if !ok {
			return false
		}
	}
	return true
}

func isCovered1(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] < ranges[j][0] {
			return true
		} else {
			return false
		}
	})
	a := [][]int{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] > a[len(a)-1][1]+1 {
			a = append(a, ranges[i])
		} else {
			a[len(a)-1][1] = max(ranges[i][1], a[len(a)-1][1])
		}
	}

	for _, v := range a {
		if left >= v[0] && right <= v[1] {
			return true
		}
	}
	return false
}

func main() {
	a := [][]int{[]int{1, 1}}
	fmt.Println(isCovered(a, 1, 50))
}
