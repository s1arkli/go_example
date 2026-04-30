package main

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 其实就是中序遍历然后到第k个的时候返回
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	if root == nil {
		return 0
	}

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		k--
		if k == 0 {
			res = root.Val
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}
