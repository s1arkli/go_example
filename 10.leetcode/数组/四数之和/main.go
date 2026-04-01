package main

import (
	"sort"
)

/*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)

	//和三数之和类似，只是加一层for，多固定一个值
	for i := 0; i < n-3; i++ {
		//本题目还需要不重复的答案，由于数组已经排序，所以只需要跳过相同的值即可保证答案是唯一的。
		//还需要跳过第一个保证每种数字至少使用过一遍，与前一位比较就可以保证这个数字至少参与过一次，跳过即可
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			//同理
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]

				if sum == target {
					re := []int{nums[i], nums[j], nums[l], nums[r]}
					result = append(result, re)

					for l < r && nums[l] == nums[l+1] {
						l++
					}

					for l < r && nums[r] == nums[r-1] {
						r--
					}
					//收缩
					l++
					r--
				} else if sum < target {
					l++
				} else if sum > target {
					r--
				}
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
