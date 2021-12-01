package main

import "fmt"

var count int

func uniquePaths(m int, n int) int {
	count = 0
	dfsss(0, 0, m, n)
	return count
}

// m ,n 过大会超时
func dfsss(i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	} else if i == m-1 && j == n-1 {
		count = count + 1
	} else {
		dfsss(i+1, j, m, n)
		dfsss(i, j+1, m, n)
	}
}

func dongtaiguihua(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m-1; i++ {
		dp[i][n-1] = 1
	}
	for j := 0; j < n-1; j++ {
		dp[m-1][j] = 1
	}
	dp[m-1][n-1] = 0
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}

func main() {
	fmt.Println(dongtaiguihua(23, 12))
}
