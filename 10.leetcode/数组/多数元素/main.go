package main

import "sort"

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
*/

// 哈希、排序、Boyer-Moore 投票算法
func majorityElement_(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement(nums []int) int {
	a, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			a = num
		}

		if num == a {
			count++
		} else {
			count--
		}
	}
	return a
}
