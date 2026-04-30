package main

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	res := make([]*TreeNode, 0)

	var a func(root *TreeNode)
	a = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root)
		a(root.Left)
		a(root.Right)
	}
	a(root)

	n := len(res)
	for i := 0; i < n; i++ {
		if i == n-1 {
			res[i].Left, res[i].Right = nil, nil
			break
		}
		res[i].Left = nil
		res[i].Right = res[i+1]
	}
}
