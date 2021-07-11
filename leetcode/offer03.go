package main

func findRepeatNumber(nums []int) int {
	a := make(map[int]int)
	for i := 0;  i < len(nums); i++ {
		if  _ ,bool := a[nums[i]]; bool {
			return nums[i]
		} else {
			a[nums[i]] = 1
		}
	}
	return -1
}
