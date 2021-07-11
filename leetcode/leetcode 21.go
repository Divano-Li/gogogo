package main

import "fmt"

func main() {
	fmt.Println(1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	temp := l2
	if l1.Val >= l2.Val {
		temp = l2
		temp.Next = mergeTwoLists(l1, l2.Next)
	}

	if l1.Val < l2.Val {
		temp = l1
		temp.Next = mergeTwoLists(l1.Next, l2)
	}
	return temp
}
