package main

/*给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。*/

func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)
	res := make([]int, 0)
	for _, num := range nums1 {
		seen[num] = true
	}
	for _, num := range nums2 {
		if seen[num] {
			delete(seen, num)
			res = append(res, num)
		}
	}
	return res
}
