package main

import "fmt"

// 实现一个队列

func main() {
	res := make([][]int, 0)
	//a := make()
	if t:=res[0]; t!=nil{
		fmt.Print(1)
	}
	fmt.Print(res[0])
}

type Tree struct {
	tree *TreeNode
	level int
}
type queue []*Tree

func (q *queue) getHead() *Tree {
	if len(*q) != 0 {
		res := (*q)[0]
		len := len(*q)
		*q = (*q)[1:len]
		return  res
	}
	return nil
}

func (q *queue) appendTail(value *Tree) {
	if value.tree != nil {
		*q = append(*q, value)
	}
}

func (q *queue) isEnd() bool {
	return len(*q) == 0
}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	a := queue{}
	a.appendTail(&Tree{root,0})
    res := make([][]int, 0)
    lever := 0
    end := a.isEnd()
    for temp := a.getHead() ; !end ; temp = a.getHead() {
    	if len(res)  < temp.level +1 {
    		res = append(res, []int{temp.tree.Val})
		} else {
			res[lever] = append(res[lever], temp.tree.Val)
		}
    	a.appendTail(&Tree{temp.tree.Left, temp.level + 1})
		a.appendTail(&Tree{temp.tree.Right, temp.level + 1})
    	end = a.isEnd()
	}
	return res
}

//方法一的解法过于复杂, 还专门构造了 queue struct, 换一种方式
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int,0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	level :=0
	for ; len(queue) !=0; {
		res = append(res, []int{})
		temp := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right !=nil {
				temp = append(temp, v.Right)
			}
			res[level] = append(res[level],v.Val)
		}
		level++
		queue = temp
	}
	return res
}


func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int,0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	level :=0
	for ; len(queue) !=0; {
		res = append(res, []int{})
		temp := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right !=nil {
				temp = append(temp, v.Right)
			}
			if level % 2 == 0 {
				res[level] = append(res[level],v.Val)
			}
		}
		if level % 2 != 0 {
			for i := len(queue)-1; i >=0 ; i-- {
				res[level] = append(res[level],queue[i].Val)
			}
		}
		level++
		queue = temp
	}
	return res
}