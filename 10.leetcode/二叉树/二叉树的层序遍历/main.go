package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	queue = append(queue, root)

	for n := len(queue); n > 0; {

		level := make([]int, 0)
		for i := 0; i < n; i++ {
			node := queue[i]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
		queue = queue[n:]
		n = len(queue)
	}
	return res
}
