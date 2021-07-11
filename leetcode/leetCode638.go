package main

import (
	"math"
)

func judgeSquareSum(c int) bool {
	maxInt := math.Ceil(math.Sqrt(float64(c)))
	m := map[float64]int{}
	var i float64
	i = 0
	for ; i <= maxInt; i++ {
		m[math.Pow(i, 2)] = 1
	}
	temp := float64(c)
	for j := range m {
		if m[temp-j] == 1 {
			return true
		}
	}
	return false
}

func main() {
	judgeSquareSum(1)
}
