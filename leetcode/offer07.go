package main

//平衡二叉树的话， 栈的深度为logN, 时间复杂度为：   空间复杂度：
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) ==0 {
		return nil
	}
	first := preorder[0]
	preorder = preorder[1:len(preorder)]

	index :=0
	for i,v :=range inorder {
		if v == first {
			index = i
			break
		}
	}
	res := TreeNode{Val: first}
	inLeft := inorder[0:index]
	inRight := inorder[index+1:]



	lenInLeft := len(inLeft)
	preLeft := preorder[0:lenInLeft]
	preRight := preorder[lenInLeft:]

	res.Left = buildTree(preLeft, inLeft)
	res.Right = buildTree(preRight, inRight)

	return &res
}
