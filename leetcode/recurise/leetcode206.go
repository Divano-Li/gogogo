package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	} else {
		/*temp := reverseList(head.Next)
		temp.Next = head
		head.Next = temp
		return head*/  //这种写法 不能倒序， 所以，我们应该想到 返回的时temp

		temp := reverseList(head.Next)
		head.Next.Next  = head
		head.Next = nil
		return temp
	}

}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
    var t *ListNode
    var temp *ListNode
    temp = nil
	for ;head != nil;  {
		 t = &ListNode{head.Val, nil}
		 t.Next = temp
		 temp = &ListNode{t.Val, t.Next}
		 head = head.Next
	}
	return t

}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode = nil
	var current = head
	for ; current != nil; {
		temp := current.Next
		current.Next = pre
		pre = current
		current = temp
	}
	return pre
}