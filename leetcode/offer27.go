package main

func mirrorTree(root * TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := &TreeNode{root.Val, nil, nil}
	res.Left = mirrorTree(root.Right)
	res.Right = mirrorTree(root.Left)
	return res
}
