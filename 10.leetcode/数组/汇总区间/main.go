package main

import "fmt"

/*
给定一个  无重复元素 的 有序 整数数组 nums 。

区间 [a,b] 是从 a 到 b（包含）的所有整数的集合。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个区间但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b
*/

func summaryRanges(nums []int) (ans []string) {
	for i := 0; i < len(nums); {
		left := i

		for i++; i < len(nums) && nums[i] == nums[i-1]+1; {
			i++
		}

		if i-left == 1 {
			ans = append(ans, fmt.Sprintf("%d", nums[left]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[left], nums[i-1]))
		}
	}
	return
}
