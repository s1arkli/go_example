package main

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。*/

func productExceptSelf(nums []int) []int {
	//基本思想就是左边的乘积乘右边的乘积，可以参考前缀和和思想，弄一个前缀积和后缀积。
	left, right := make([]int, len(nums)+1), make([]int, len(nums)+1) //+1是因为初始值应设置为1，遍历时直接做乘法即可
	left[0] = 1
	right[len(nums)] = 1
	res := make([]int, len(nums))

	//两次遍历得到前缀后缀积
	for i := 0; i < len(nums); i++ {
		left[i+1] = left[i] * nums[i]
	}

	for i := len(nums); i > 0; i-- {
		right[i-1] = right[i] * nums[i-1]
	}

	//注意边界问题
	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i+1]
	}
	return res
}
