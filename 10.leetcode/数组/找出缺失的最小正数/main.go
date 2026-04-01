package main

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
*/

// firstMissingPositive1 不考虑空间，可以使用map保存nums元素，从1开始遍历，不存在的return
func firstMissingPositive1(nums []int) int {
	seen := make(map[int]bool)
	maxNum := 0
	for _, n := range nums {
		if n > maxNum {
			maxNum = n
		}
		seen[n] = true
	}

	for i := 1; i <= maxNum; i++ {
		if !seen[i] {
			return i
		}
	}
	return maxNum + 1
}

// firstMissingPositive2 由于要求空间大小，所以只能在原地进行交换。
func firstMissingPositive2(nums []int) int {
	//参考别人的思路，大致就是把元素和下标对应上。再次从1开始遍历，一旦nums[i] != i+1，即证明i+1就是缺失的最小正数。因为原数组是从1开始排的
	//（下标从0开始）
	for i := 0; i < len(nums); i++ {
		//细节：只有1开始有排序的意义，所以nums[i]>0;如果大于len(nums)，证明从1开始就没法连续到len(nums)，所以限定小于等于len;最后是防止死循环，
		//当有两个相同元素时，只需要一个在位置上即可。;循环放置，保证每一个1-n的数字都在正确的位置上。
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			//交换放到正确的位置上
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
