package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
  left := maxDepth(root.Left)
  right := maxDepth(root.Right)
  if left >=  right {
  	return left + 1
  } else {
  	return right + 1
  }
}

func maxDepth1(root *TreeNode) int {
	temp := []*TreeNode{root}
	if root == nil {
		return 0
	}
	temp = append(temp, root)
	temp1 := []*TreeNode{}
	time := 0
	for ;len(temp)!=0; {
		for i:=0; i< len(temp); i++ {
			if temp[i].Left != nil {
				temp1 = append(temp1,temp[i].Left)
			}
			if temp[i].Right != nil {
				temp1 = append(temp1, temp[i].Right)
			}
		}
		temp = temp1
		temp1 = nil
		time++
	}
	return time
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := 0
	temp := []*TreeNode{root}
	for ; len(temp) != 0 ; {
		len := len(temp)
		for ;len !=0;len-- {
			if temp[0].Left != nil {
				temp = append(temp, temp[0].Left)
			}
			if temp[0].Right != nil {
				temp = append(temp, temp[0].Right)
			}
			temp = temp[1:]
		}
		l++
	}
	return l
}



func main() {
	/*a :=  &TreeNode{0, nil, nil}
	b := maxDepth1(a)
	fmt.Println(b)*/
	a := []int{1,2}
	b := a[0:0]
	fmt.Println(b)
}