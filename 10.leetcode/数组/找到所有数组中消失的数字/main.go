package main

/*
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
*/

func findDisappearedNumbers(nums []int) []int {
	//使用下标进行对应，在原数组上进行标记
	res := make([]int, 0)
	for _, v := range nums {
		idx := abs(v) - 1

		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
