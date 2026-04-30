package main

/*
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 计算直径就是找每个节点最长的左右节点之和，也就是左右节点的深度之和
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var long = 0
	deep(root, &long)
	return long - 1
}

func deep(root *TreeNode, long *int) int {

	if root == nil {
		return 0
	}
	left := deep(root.Left, long)
	right := deep(root.Right, long)
	if left+right+1 > *long {
		*long = left + right + 1
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
