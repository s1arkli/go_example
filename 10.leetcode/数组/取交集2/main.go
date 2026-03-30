package main

import "sort"

/*
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。
返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
*/

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)

	for l, r := 0, 0; l < len(nums1) && r < len(nums2); {
		n1, n2 := nums1[l], nums2[r]
		if n1 == n2 {
			res = append(res, n1)
			l++
			r++
		} else if n1 < n2 {
			l++
		} else {
			r++
		}
	}
	return res
}
