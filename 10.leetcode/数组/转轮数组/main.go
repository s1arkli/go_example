package main

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/

func rotate(nums []int, k int) {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < len(nums); i++ {
		res[index(n, i, k)] = nums[i]
	}
	copy(nums, res)
}

func index(n, a, k int) int {

	idx := a + (k % n)
	if idx >= n {
		return idx - n
	}
	return idx
}
