package main

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)
	var backStack func(start int)

	backStack = func(start int) {
		if start == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backStack(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	backStack(0)
	return res
}
