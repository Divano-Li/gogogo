package main

import "fmt"



type TreeNode struct {
 	Val int
    Left *TreeNode
    Right *TreeNode
  }


func main() {
	a := generateTrees(3)
	fmt.Println(len(a))
}

func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{1, nil, nil}}
	}
	return  generates(1,n)

}

func generates(start int, end int) []*TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return []*TreeNode{&TreeNode{start, nil, nil}}
	}
	result := make([]*TreeNode, 0)
	for i:= start; i<= end ; i++ {

		left := generates(start, i-1)
		right := generates(i+1, end)
		if left != nil && right != nil {
			for j:=0 ; j<len(left); j++ {
				for k:=0; k< len(right); k++ {
					temp := &TreeNode{i, nil, nil}
					temp.Left = left[j]
					temp.Right = right[k]
					result = append(result, temp)
				}
			}

		} else if left != nil && right == nil {
			for j:=0; j<len(left); j++ {
				temp := &TreeNode{i, nil, nil}
				temp.Left = left[j]
				temp.Right = nil
				result = append(result, temp)
			}

		}else if right != nil && left == nil {
			for j:=0; j<len(right); j++ {
				temp := &TreeNode{i, nil, nil}
				temp.Right = right[j]
				temp.Left = nil
				result = append(result, temp)
			}
		} else {
			temp := &TreeNode{i, nil, nil}
			temp.Right = nil
			temp.Left = nil
			result = append(result, temp)
		}

	}
    return result
}