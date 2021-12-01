package main

/*
 [2,3,1,1,4]
 2+3+1+1
*/

func canJump(nums []int) bool {
	l := len(nums)
	max := 0
	for i := 0; i < len(nums) && i <= max; i++ {
		if i+nums[i] >= l-1 {
			return true
		} else {
			max = maxInt(max, i+nums[i])
		}
	}
	return false
}
