package main

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/

func maxSlidingWindow(nums []int, k int) []int {
	//维护一个窗口，内部的队头也就是最左侧永远是最大的，新进来的更大则踢掉所有值，更小则跟在后面，这样每次进来只需要和队头进行比较。
	queue := make([]int, k)
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}

		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)
		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}
