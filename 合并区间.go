package main

import "sort"

func merge(interval [][]int) [][]int {
	sort.Slice(interval, func(i, j int) bool {
		if interval[i][0] < interval[j][0] {
			return true
		} else {
			return false
		}
	})
	result := make([][]int, 0)
	for _, v := range interval {
		temp := v
		if len(result) == 0 {
			result = append(result, temp)
		} else {
			pre := result[len(result)-1]
			if temp[0] > pre[1] {
				result = append(result, temp)
			} else {
				result[len(result)-1][1] = max(temp[1], pre[1])
			}
		}
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
