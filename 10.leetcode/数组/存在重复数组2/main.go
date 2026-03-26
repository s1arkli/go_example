package main

/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。
如果存在，返回 true ；否则，返回 false 。
*/

// 哈希
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for idx, num := range nums {
		if v, ok := m[num]; ok {
			if idx-v <= k {
				return true
			}
		}
		m[num] = idx
	}
	return false
}
