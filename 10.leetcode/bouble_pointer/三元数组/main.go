package main

import (
	"slices"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	slices.Sort(nums)

	//-2是因为三元数组，至少要3个不同int，所以保留最后两位
	for i := 0; i < len(nums)-2; i++ {
		//按照从小到大排序，如果开头>0，则后续之和只会>0
		if nums[i] > 0 {
			break
		}

		//跳过开头数相同的哪些
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}
