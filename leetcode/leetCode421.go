package main

import "fmt"

type PreTree struct {
	Left  *PreTree
	Right *PreTree
}

func findMaximumXOR(nums []int) int {
	// 第一步是前缀树, 该前缀树是一个二叉树， 0,1 表示，高位和地位
	preTree := &PreTree{}
	for i := 0; i < len(nums); i++ {
		insert(nums[i], preTree)
	}
	// 第二步， 怎么找到最大值呢？  贪心算法， 每一层找到最大值
	ress := 0
	for i := 0; i < len(nums); i++ {
		root := preTree
		total := 0
		for j := 31; j >= 0; j-- {
			if (nums[i]>>j)&1 == 1 {
				if root.Left != nil {
					total = total*2 + 1
					root = root.Left
				} else {
					total = total * 2
					root = root.Right
				}
			} else {
				if root.Right != nil {
					total = total*2 + 1
					root = root.Right
				} else {
					total = total * 2
					root = root.Left
				}
			}
		}
		ress = maxs(ress, total)
	}
	return ress
}

func maxs(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func insert(a int, tree *PreTree) {
	res := make([]int, 32)
	res[0] = a & 1
	for i := 1; i < 32; i++ {
		a = a >> 1
		res[i] = a & 1
	}
	for i := len(res) - 1; i >= 0; i-- {
		if res[i]&1 == 1 {
			if tree.Right == nil {
				tree.Right = &PreTree{}
			}
			tree = tree.Right
		} else {
			if tree.Left == nil {
				tree.Left = &PreTree{}
			}
			tree = tree.Left
		}
	}
}

func main() {
	a := 0
	fmt.Print(a >> 1)
}
