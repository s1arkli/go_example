package main

/*
给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。

换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。

以数组形式返回答案。
*/

// 暴力双循环解法
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[j] < nums[i] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

// 学习到的解法
func smallerNumbersThanCurrent2(nums []int) []int {
	//由于限制了输入数组的值域，可以把数组内的所有数出现的次数统计下来（使用下标对应），然后在下标左侧所有数的和即是小于该数的数字的数目。
	idxNum := [101]int{}

	for _, v := range nums {
		//统计每个数字出现的次数
		idxNum[v]++
	}

	//求前缀和
	for i := 1; i < 101; i++ {
		idxNum[i] += idxNum[i-1]
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		//0不存在更小的数了
		if nums[i] == 0 {
			res = append(res, 0)
		} else {
			res = append(res, idxNum[nums[i]-1])
		}
	}
	return res
}
