package main

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		third := len(nums) - 1
		for second := i + 1; second < len(nums); second++ {
			if second > 1 && nums[second] == nums[second-1] {
				continue
			}
			for third > second && nums[third]+nums[second]+nums[i] > 0 {
				third--
			}
			if third == second {
				break
			}
			if nums[third]+nums[second]+nums[i] == 0 {
				result = append(result, []int{nums[i], nums[second], nums[third]})
			}
		}
	}
	return result
}

//双指针的问题用 while 循环比较好
func threeNum1(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		target := -nums[i]
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
