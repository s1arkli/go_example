package main

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2
*/

func subarraySum(nums []int, k int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		h := nums[i]

		if h == k {
			sum++
		} else {
			for j := i + 1; j < len(nums); j++ {
				if h += nums[j]; h == k {
					sum++
				}
			}
		}
	}
	return sum
}

func subarraySum1(nums []int, k int) int {
	prefixSum := make(map[int]int, len(nums))
	prefixSum[0] = 1
	count := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if cnt, ok := prefixSum[sum-k]; ok {
			count += cnt
		}
		prefixSum[sum]++
	}

	return count
}
