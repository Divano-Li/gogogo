package main


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type stack1 []*TreeNode

func (this *stack1) push(value *TreeNode) {
	*this = append(*this, value)
}

func (s *stack1) pop() *TreeNode {
	if len(*s) == 0 {
		return nil
	}
	len := len(*s)
	 *s = (*s)[:len-1]
	 return (*s)[len-1]
}

type Queue1 []*TreeNode

func (q *Queue) getHead() *TreeNode{
	if len(*q) == 0 {
		return nil
	}
	len := len(*q)
	res := (*q)[0]
	*q = (*q)[1:len]
	return res
}

func (q *Queue) appendTail(value *TreeNode){
	*q = append(*q,value)
}

func (q *Queue) length() int {
	return len(*q)
}

// 时间复杂度 需要遍历n个节点， 空间复杂度 如果是平衡二叉树，
// 那么的话每次add 两个节点，出去一个节点， 最多存在n/2 个节点在队列中， 所以需要O(n)的空间
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := Queue{}
	res := make([]int,0)
	queue.appendTail(root)
	length := 1
	for temp:= queue.getHead(); length !=0 ; temp= queue.getHead(){
		if temp != nil {
			res = append(res, temp.Val)
			queue.appendTail(temp.Left)
			queue.appendTail(temp.Right)
		}
		length = queue.length()
	}
	return res
}
