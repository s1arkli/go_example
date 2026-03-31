package main

import "sort"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(intervals [][]int) (ans [][]int) {
	//该题目优先排序，然后遍历进行区间合并
	if len(intervals) == 0 {
		return
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	end := ans[0][1]

	for _, v := range intervals {
		if v[0] > end {
			ans = append(ans, v)
			end = v[1]
		} else {
			bigger := max(end, v[1])
			ans[len(ans)-1][1] = bigger
			end = bigger
		}
	}
	return
}
