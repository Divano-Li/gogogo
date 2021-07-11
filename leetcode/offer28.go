package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node32 := &TreeNode{
		2,nil,nil,
	}
	node31 := &TreeNode{
		2, nil, nil,
	}
	node22 := &TreeNode{
		2,  nil ,node32,
	}
	node21 := &TreeNode{
		2,node31,node31,
	}
	node1 := &TreeNode{
		1,node21, node22,
	}
	res := isSymmetric(node1)
	fmt.Print(res)
}

func isSymmetric1(root *TreeNode) bool {
   if root == nil {
   	 return true
   }
   return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right ==nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 二叉树 左根右 的遍历方式跟 右根左的遍历方式 结果相同，那么就能判断对称的了
// 这种方式 遍历会有bug  1 2 2 2 null 2  不对称，但是该方法会判断为对称
func isSymmetric(root *TreeNode) bool {
	// 左根右的遍历
	a := make([]int, 0)
	b := make([]int, 0)
	recurise1(root, &a)
	recurise2(root, &b)
	for i:= range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func recurise1(node *TreeNode, result *[]int) {
	if node != nil {
		recurise1(node.Left, result)
		*result = append(*result, node.Val)
		recurise1(node.Right, result)
	}
}

func recurise2(node *TreeNode, result *[]int) {
	if node != nil {
		recurise2(node.Right, result)
		*result = append(*result, node.Val)
		recurise2(node.Left, result)
	}
}
