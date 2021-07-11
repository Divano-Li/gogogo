package main

import "sort"

func merge(intervals [][]int) [][]int {
	results := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] - intervals[j][0] < 0
	})
	results = append(results, intervals[0])
	for i:= 1; i< len(intervals);i++ {
		last := results[len(results)-1]
		if intervals[i][0] > last[1] {
			results = append(results, intervals[i])
		} else {
			results[len(results)-1][1] = maxInt(intervals[i][1], last[1])
		}
	}
	return results
}

func maxInt(x, y int) int {
	if x>y {
		return x
	} else {
		return y
	}
}
