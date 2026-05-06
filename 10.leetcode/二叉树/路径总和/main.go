package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
暴力dfs，在每个节点都去深度优先进行检索
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	count := 0

	if root.Val == targetSum {
		count++
	}

	count += dfs(root.Left, targetSum-root.Val)
	count += dfs(root.Right, targetSum-root.Val)
	return count
}
*/

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	prefix := make(map[int]int)
	prefix = map[int]int{0: 1}
	return a(root, prefix, targetSum, 0)
}

func a(root *TreeNode, prefix map[int]int, target, preSum int) int {
	if root == nil {
		return 0
	}

	currSum := preSum + root.Val
	nowCount := 0
	v, ok := prefix[currSum-target]
	if ok {
		nowCount = v
	}
	prefix[currSum]++
	nowCount += a(root.Left, prefix, target, currSum)
	nowCount += a(root.Right, prefix, target, currSum)

	prefix[currSum]--
	return nowCount
}
