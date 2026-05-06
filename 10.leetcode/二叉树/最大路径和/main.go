package main

import "math"

/*
给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 函数的作用是返回该节点下的最大路径和
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxInt := math.MinInt32

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		maxInt = max(maxInt, left+right+node.Val)
		return node.Val + max(left, right)
	}
	dfs(root)
	return maxInt
}
