package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//1,2,3,4,5
// 回溯算法
var count = 0
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count = 0
	pre := &ListNode{0, head}
	trackback(pre, n)
	return pre.Next
}

func trackback(head *ListNode, n int) {
	if head.Next == nil {
		count = count+1
		return
	} else {
		trackback(head.Next, n)
		if count == n {
			head.Next = head.Next.Next
		}
		count = count +1
	}
}


// 双指针算法
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	pre := &ListNode{0,head}
	slow := pre
	fast := pre
	count := 0
	for count < n {
		fast = fast.Next
		count++
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return  pre.Next
}