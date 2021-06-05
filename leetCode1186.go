package main

func maxmumSum(arr []int) int {
	dp := make([][]int, len(arr))
	for i:=0; i<len(arr); i++ {
		dp[i] = []int{0,1}
	}
	dp[0][0] = arr[0]
	dp[0][1] = 0
	result := dp[0][0]
	for i:= 1; i< len(arr); i++ {
		dp[i][0] = max(arr[i], dp[i-1][0]+arr[i])
		dp[i][1] = max(dp[i-1][1] + arr[i], dp[i-1][0] - arr[i-1] + arr[i])
		result = max(result, max(dp[i][0], dp[i][1]))
	}
	return result
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}