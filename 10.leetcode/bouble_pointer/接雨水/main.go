package main

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap(height []int) int {
	sum := 0
	n := len(height)
	if n < 3 {
		return sum
	}

	ltMax := make([]int, n)
	rtMax := make([]int, n)

	ltMax[0] = height[0]
	rtMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		ltMax[i] = max(ltMax[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		rtMax[i] = max(rtMax[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		sum += min(ltMax[i], rtMax[i]) - height[i]
	}
	return sum

}
