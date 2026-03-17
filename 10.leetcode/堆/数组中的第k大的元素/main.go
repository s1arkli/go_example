package main

import (
	"fmt"
	"math/rand"
)

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}

func findKthLargest_(nums []int, k int) int {
	//快速排序，找到下标=n-k的元素
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, left, right, target int) int {
	if left == right {
		return nums[left]
	}

	pp := left + rand.Intn(right-left+1)
	ppnum := nums[pp]

	nums[pp], nums[right] = nums[right], nums[pp]

	storeIdx := left
	for i := left; i < right; i++ {
		if nums[i] < ppnum {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
	}
	nums[right], nums[storeIdx] = nums[storeIdx], nums[right]

	if storeIdx == target {
		return nums[storeIdx]
	} else if storeIdx < target {
		return quickSort(nums, storeIdx+1, right, target)
	} else {
		return quickSort(nums, left, storeIdx-1, target)
	}
}
