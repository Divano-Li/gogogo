package main

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
  dp:= make([][]int, len(matrix))
  for i:=0; i<len(dp);i++ {
  	dp[i] = make([]int, len(matrix[0]))
  }
  res := make([]int, 0)
  for i:=0;i<len(dp);i++ {
  	for j:=0; j < len(dp[i]); j++ {
  		if i ==0 && j== 0 {
  			dp[i][j] = matrix[0][0]
		} else if i == 0 {
			dp[i][j] = dp[i][j-1] ^ matrix[i][j]
		} else if j == 0 {
			dp[i][j] = dp[i-1][j] ^ matrix[i][j]
		} else {
			dp[i][j] = dp[i-1][j] ^ dp[i][j-1] ^ dp[i-1][j-1]
		}
		res = append(res, dp[i][j])
	}
  }
  sort.Ints(res)
  return res[len(res) - k]
}
