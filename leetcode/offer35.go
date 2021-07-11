package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	addressmap := make(map[*Node]*Node)
	cur := head
	for head != nil {
		temp := &Node{head.Val, nil, nil}
		addressmap[head] = temp
		head = head.Next
	}
	res, _ := addressmap[cur]
	res1 := res
	for cur != nil {
		res, _ = addressmap[cur]
		if cur.Next != nil {
			res.Next, _ = addressmap[cur.Next]
		} else {
			res.Next = nil
		}
		if cur.Random != nil {
			res.Random, _ = addressmap[cur.Random]
		} else {
			res.Random = nil
		}
		cur = cur.Next
		res = res.Next
	}
	return res1
}
