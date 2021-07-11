package main

import "fmt"

func main() {


	a := sortArray([]int{-2,3,-5})
	fmt.Println(a)
}
func sortArray(nums []int) []int {
	if nums == nil || len(nums)==0{
		return nil
	}
	len := len(nums)
	index := 0
	for i,j:= 0, len-1; i < j; {
		for ; i< j; j-- {
			if nums[j] < nums[i] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
				index = j
				break
			}
		}
		for ; i < j ; i++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
				index = i
				break
			}
		}
	}

	if index > 1 {
		left := nums[0:index]
		sortArray(left)
	}

	if index+1 < len {
		right := nums[index+1: len]
		sortArray(right)
	}
	return nums
}