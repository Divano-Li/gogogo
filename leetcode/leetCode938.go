package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum = 0

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	} else {
		if root.Val >= low && root.Val <= high {
			sum = sum + root.Val
			fmt.Println(root.Val)
		}
		rangeSumBST(root.Left, low, high)
		rangeSumBST(root.Right, low, high)
	}
	return sum
}
