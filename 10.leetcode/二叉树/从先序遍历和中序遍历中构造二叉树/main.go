package main

/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rVal := preorder[0]

	li, ri := make([]int, 0), make([]int, 0)
	leftLen := 0
	for k, v := range inorder {
		if v == rVal {
			leftLen = k
			li = inorder[:k]
			ri = inorder[k+1:]
			break
		}
	}

	lr, rr := make([]int, 0), make([]int, 0)
	lr = preorder[1 : leftLen+1]
	rr = preorder[leftLen+1:]

	root := &TreeNode{
		Val:   rVal,
		Left:  buildTree(lr, li),
		Right: buildTree(rr, ri),
	}
	return root
}
