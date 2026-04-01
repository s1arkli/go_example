package main

import "sort"

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个在 不同下标位置 的整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return sum
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
