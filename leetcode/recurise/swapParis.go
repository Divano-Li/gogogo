package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
   1 -> 2 -> 3 -> 4
   2 -> 1 -> 4 -> 3
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next ==nil{
		return head
	}

	second := head.Next
	head.Next = swapPairs(second.Next)
	second.Next = head
	return  second

	/*
	将second 去除掉
	此方法head.Next 的值已经变了， 所以不能这么做
	head.Next =swapPairs(head.Next.Next)
	head.Next.Next = head
	return  head.Next*/
}
